// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/dnsc
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

// Package dqrs provides types and functions used by this application to
// collect and process DNS queries and responses.
package dqrs

import (
	"bytes"
	"fmt"
	"net"
	"sort"
	"strings"
	"time"

	"github.com/miekg/dns"
)

// defaultDNSPort is the default UDP and TCP port used for incoming DNS
// requests
const defaultDNSPort = "53"

// DNSQueryResponse represents a query and response from a DNS server.
// Multiple records may be returned for a single query (e.g., CNAME and A
// records).
type DNSQueryResponse struct {

	// Answer may potentially be composed of multiple Resource Record types
	// such as CNAME and A records. We later separate out the types when
	// needed.
	Answer []dns.RR

	// Server is the DNS server used for this query and response.
	Server string

	// Query is the FQDN that we requested a record for.
	Query string

	// RequestedRecordType represents the type of record requested as part of
	// the query
	RequestedRecordType uint16

	// Error records whether an error occurred during any part of performing a
	// query
	QueryError error

	// ResponseTime, also known as the Round-trip Time, can be best summed up
	// by this Cloudflare definition: "Round-trip time (RTT) is the duration
	// in milliseconds (ms) it takes for a network request to go from a
	// starting point to a destination and back again to the starting point."
	ResponseTime time.Duration
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

// Records returns a comma-separated string of all DNS records retrieved by an
// earlier query. The output is formatted for display in a Tabwriter table.
func (dqr DNSQueryResponse) Records() string {

	records := make([]string, 0, 5)

	for _, record := range dqr.Answer {

		var answer string

		// FIXME: How to dynamically get a "short" string value for each
		// record type so that we don't have to hard-code in a switch
		// statement and then use a type-specific field or method to retrieve
		// a text copy of the value? For example, *dns.CNAME type requires
		// use of v.Target (field value) to get a usable string, whereas
		// v.AAAA type has a usable String() method.

		switch v := record.(type) {
		case *dns.A:
			answer = v.A.String() + " (A)"
		case *dns.AAAA:
			answer = v.AAAA.String() + " (AAAA)"
		case *dns.CNAME:
			answer = v.Target + " (CNAME)"
			//fmt.Println("Check *dns.CNAME type switch case stmt")
			// answer = v.String() + " (CNAME)"
		case *dns.MX:
			answer = v.Mx + " (MX)"
		default:
			answer = "type unknown"
		}

		records = append(records, answer)
	}

	return strings.Join(records, ", ")
}

// TTLs returns a comma-separated list of record TTLs from an earlier query
func (dqr DNSQueryResponse) TTLs() string {

	ttlEntries := make([]string, 0, 5)

	for _, record := range dqr.Answer {
		ttlEntries = append(ttlEntries, fmt.Sprint(record.Header().Ttl))
	}

	return strings.Join(ttlEntries, ", ")
}

// PerformQuery wraps the bulk of the query/record logic performed by this
// application
func PerformQuery(query string, server string, qType uint16, timeout time.Duration) DNSQueryResponse {

	var msg dns.Msg

	fqdn := dns.Fqdn(query)

	// NOTE: Recursion is used by default, which resolves CNAME entries
	// back to the actual A or AAAA records
	msg.SetQuestion(fqdn, qType)

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
		dnsQueryResponse.QueryError = fmt.Errorf("no records found for query")
		return dnsQueryResponse
	}

	dnsQueryResponse.Answer = in.Answer

	return dnsQueryResponse
}
