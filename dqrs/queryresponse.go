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
	"sort"
	"strings"
	"unicode"

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

// SortRecords sorts DNS query responses by the response value so that like
// answers from multiple servers are more likely to be aligned visually
func (dqr *DNSQueryResponse) SortRecords() {

	sort.Slice(dqr.Answer, func(i, j int) bool {

		var indexI string

		switch v := dqr.Answer[i].(type) {
		case *dns.A:
			indexI = v.A.String()
		case *dns.CNAME:
			indexI = v.Target
		default:
			indexI = "type unknown"
		}

		var indexJ string
		switch v := dqr.Answer[j].(type) {
		case *dns.A:
			indexJ = v.A.String()
		case *dns.CNAME:
			indexJ = v.Target
		default:
			indexJ = "type unknown"
		}

		// force CNAME types to the front of the line
		fmt.Printf("Is indexI (%s) less than indexJ (%s)? %t\n", indexI, indexJ, indexI < indexJ)
		switch {
		case unicode.IsLetter(rune(indexI[0])) && unicode.IsLetter(rune(indexJ[0])):
			fmt.Printf("Both start with letters: %v, %v\n", indexI, indexJ)
			return indexI < indexJ

		case unicode.IsLetter(rune(indexI[0])) && !unicode.IsLetter(rune(indexJ[0])):
			fmt.Printf("Only indexI starts with a letter: %v, %v\n", indexI, indexJ)
			return false

		case !unicode.IsLetter(rune(indexI[0])) && unicode.IsLetter(rune(indexJ[0])):
			fmt.Printf("Only indexJ starts with a letter: %v, %v\n", indexI, indexJ)
			return true

		default:
			fmt.Printf("Both do not start with letters: %v, %v\n", indexI, indexJ)
			return indexI < indexJ
		}

		// if unicode.IsLetter(rune(indexI[0])) {
		// 	// testRune := rune('a')
		// 	// if unicode.IsLetter(testRune) {
		// 	fmt.Printf("First character is a letter: %v\n", rune(indexI[0]))
		// 	// fmt.Printf("First character is a letter: %v\n", testRune)
		// 	return true
		// }
		// return indexI < indexJ

		return indexI < indexJ

	})

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

	// Record the reliable DNS-related details we have thus far. Use zero
	// value initially for Answer field. We'll set a value for QueryError if
	// needed later.
	dnsQueryResponse := DNSQueryResponse{
		Server: server,
		Query:  query,
	}

	// Perform UDP-based query using default settings
	in, err := dns.Exchange(&msg, server+":53")
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
