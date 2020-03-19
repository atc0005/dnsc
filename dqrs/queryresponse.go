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
	"fmt"
	"strings"

	"github.com/miekg/dns"
)

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

	// Error records whether an error occurred during any part of performing a
	// query
	QueryError error
}

// DNSQueryResponses is a collection of DNS query responses. Intended for
// aggregation before bulk processing of some kind.
type DNSQueryResponses []DNSQueryResponse

// Error satisfies the Error interface
// TODO: This doesn't look right
func (dqr DNSQueryResponse) Error() string {
	return fmt.Sprintf("%v", dqr.QueryError)
}

// Records returns a comma-separated string of all DNS records retrieved by an
// earlier query. The output is formatted for display in a Tabwriter table.
func (dqr DNSQueryResponse) Records() string {

	records := make([]string, 0, 5)

	for _, record := range dqr.Answer {

		var answer string

		switch v := record.(type) {
		case *dns.A:
			answer = v.A.String() + " (A)"
		case *dns.CNAME:
			answer = v.Target + " (CNAME)"
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
func PerformQuery(query string, server string) DNSQueryResponse {

	var msg dns.Msg

	fqdn := dns.Fqdn(query)

	// NOTE: Recursion is used by default, which resolves CNAME entries
	// back to the actual A record
	msg.SetQuestion(fqdn, dns.TypeA)

	// TODO: Use concurrency to query all DNS servers in bulk instead of
	// waiting for each to complete before querying the next one

	// Perform UDP-based query using default settings
	in, err := dns.Exchange(&msg, server+":53")
	if err != nil {
		// panic(err)
		return DNSQueryResponse{
			QueryError: err,
		}
	}

	// Early exit if one of the DNS servers returns an unexpected result
	if len(in.Answer) < 1 {

		return DNSQueryResponse{
			QueryError: fmt.Errorf("no records for %q from %s", query, server),
		}

	}

	dnsQueryResponse := DNSQueryResponse{
		// use zero value initially for Answer field
		Answer: in.Answer,
		Server: server,
		Query:  query,
	}

	return dnsQueryResponse
}
