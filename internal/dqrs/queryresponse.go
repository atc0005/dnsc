// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/dnsc
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package dqrs

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"sort"
	"time"

	"github.com/miekg/dns"
)

// ErrNoRecordsFound indicates that a nameserver found no records for the
// specified query.
var ErrNoRecordsFound = errors.New("no records found for query")

// DNSQueryResponse represents a query and response from a DNS server.
// Multiple records may be returned for a single query (e.g., CNAME and A
// records).
type DNSQueryResponse struct {

	// Server is the DNS server used for this query and response.
	Server string

	// Query is the FQDN that we requested a record for.
	Query string

	// Error records whether an error occurred during any part of performing a
	// query
	QueryError error

	// Answer may potentially be composed of multiple Resource Record types
	// such as CNAME and A records. We later separate out the types when
	// needed.
	Answer []dns.RR

	// ResponseTime, also known as the Round-trip Time, can be best summed up
	// by this Cloudflare definition: "Round-trip time (RTT) is the duration
	// in milliseconds (ms) it takes for a network request to go from a
	// starting point to a destination and back again to the starting point."
	ResponseTime time.Duration

	// RequestedRecordType represents the type of record requested as part of
	// the query
	RequestedRecordType uint16
}

// DNSQueryResponses is a collection of DNS query responses. Intended for
// aggregation before bulk processing of some kind.
type DNSQueryResponses []DNSQueryResponse

// Error satisfies the Error interface
// TODO: This doesn't look right
func (dqr DNSQueryResponse) Error() string {
	return fmt.Sprintf("%v", dqr.QueryError)
}

// Less compares records and indicates whether the first argument is less than
// the second argument. Preference is given to CNAME records.
func (dqr *DNSQueryResponse) Less(i, j int) bool {

	var indexI net.IP

	switch v := dqr.Answer[i].(type) {
	case *dns.A:
		indexI = v.A
	case *dns.AAAA:
		indexI = v.AAAA
	case *dns.CNAME:
		indexI = nil
	}

	var indexJ net.IP
	switch v := dqr.Answer[j].(type) {
	case *dns.A:
		indexJ = v.A
	case *dns.AAAA:
		indexI = v.AAAA
	case *dns.CNAME:
		indexJ = nil
	}

	return bytes.Compare(indexI, indexJ) < 0

}

// SortRecordsAsc sorts DNS query responses by the response value in ascending
// order with CNAME records listed first.
func (dqr *DNSQueryResponse) SortRecordsAsc() {

	// Place CNAME entries first, sort IP Addresses after
	sort.Slice(dqr.Answer, dqr.Less)

}

// SortRecordsDesc sorts DNS query responses by the response value in
// descending order with CNAME records listed last.
func (dqr *DNSQueryResponse) SortRecordsDesc() {

	// Place CNAME entries first, sort IP Addresses after
	sort.Slice(dqr.Answer, func(i, j int) bool {
		return !dqr.Less(i, j)
	})

}

// Records returns all DNS records associated with a query response.
func (dqr DNSQueryResponse) Records() []DNSRecord {

	records := make([]DNSRecord, 0, len(dqr.Answer))

	for _, record := range dqr.Answer {

		var recordVal string
		var recordType string

		// FIXME: How to dynamically get a "short" string value for each
		// record type so that we don't have to hard-code in a switch
		// statement and then use a type-specific field or method to retrieve
		// a text copy of the value? For example, *dns.CNAME type requires
		// use of v.Target (field value) to get a usable string, whereas
		// v.AAAA type has a usable String() method.

		switch v := record.(type) {
		case *dns.A:
			recordVal = v.A.String()
			recordType = RequestTypeA
		case *dns.AAAA:
			recordVal = v.AAAA.String()
			recordType = RequestTypeAAAA
		case *dns.CNAME:
			recordVal = v.Target
			recordType = RequestTypeCNAME
		case *dns.MX:
			recordVal = v.Mx
			recordType = RequestTypeMX
		case *dns.NS:
			recordVal = v.Ns
			recordType = RequestTypeNS
		case *dns.PTR:
			recordVal = v.Ptr
			recordType = RequestTypePTR
		case *dns.SRV:
			recordVal = v.Target
			recordType = RequestTypeSRV
		default:
			recordVal = recordValueUnknown
			recordType = RequestTypeUnknown
		}

		ttl := record.Header().Ttl

		records = append(records, DNSRecord{
			Value: recordVal,
			Type:  recordType,
			TTL:   ttl,
		})
	}

	return records
}

// RecordsFound indicates whether any query responses indicate records were
// found.
func (dqrs DNSQueryResponses) RecordsFound() bool {

	for i := range dqrs {

		// If a lookup failure occurs, or if there are no records found for a
		// query this field will be set. As long as at least one
		// DNSQueryResponse does not have an associated error, there is a
		// valid record to report.
		if dqrs[i].QueryError == nil {
			return true
		}
	}

	return false

}

// PerformQuery wraps the bulk of the query/record logic performed by this
// application
func PerformQuery(query string, server string, qType uint16, timeout time.Duration) DNSQueryResponse {

	var msg dns.Msg

	var qualifiedQuery string
	switch {
	case qType == dns.TypePTR:
		arpa, err := dns.ReverseAddr(query)
		if err != nil {
			return DNSQueryResponse{
				Server:              server,
				Query:               query,
				RequestedRecordType: qType,
				QueryError:          err,
			}
		}
		qualifiedQuery = arpa
	default:
		qualifiedQuery = dns.Fqdn(query)
	}

	// NOTE: Recursion is used by default. This results in CNAME entries
	// resolving back to the actual A or AAAA records.
	msg.SetQuestion(qualifiedQuery, qType)

	// Record the reliable DNS-related details we have thus far. Use zero
	// value initially for Answer field. We'll set a value for QueryError if
	// needed later.
	dnsQueryResponse := DNSQueryResponse{
		Server:              server,
		Query:               query,
		RequestedRecordType: qType,
	}

	// construct client so that we are able to override default settings
	client := dns.Client{
		Net:     "udp",
		Timeout: timeout,
	}

	// Perform UDP-based query using custom client settings
	remoteAddress := net.JoinHostPort(server, defaultDNSPort)
	in, rtt, err := client.Exchange(&msg, remoteAddress)
	dnsQueryResponse.ResponseTime = rtt
	if err != nil {
		dnsQueryResponse.QueryError = err
		return dnsQueryResponse
	}

	// Early exit if the DNS server returns an unexpected result
	if len(in.Answer) < 1 {
		dnsQueryResponse.QueryError = ErrNoRecordsFound
		return dnsQueryResponse
	}

	dnsQueryResponse.Answer = in.Answer

	return dnsQueryResponse
}
