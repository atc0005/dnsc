// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/dnsc
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/miekg/dns"

	"github.com/atc0005/dnsc/config"
)

// DNSQueryResponse represents a query and response from a DNS server
type DNSQueryResponse struct {

	// Depending on the system, there *could* be multiple A records returned
	// ARecords []*dns.A
	Answer []dns.RR
	Server string
	Query  string
}

// Records returns a comma-separated string of all DNS records retrieved by an
// earlier query. The output is formatted for display in a Tabwriter table.
func (dqr DNSQueryResponse) Records() string {

	var records []string

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

	var ttlEntries []string
	for _, record := range dqr.Answer {
		ttlEntries = append(ttlEntries, fmt.Sprint(record.Header().Ttl))
	}

	return strings.Join(ttlEntries, ", ")
}

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fqdn := dns.Fqdn(cfg.Query)

	var results []DNSQueryResponse

	// loop over each of our DNS servers, build up a results set
	for _, server := range cfg.Servers {

		var msg dns.Msg

		// NOTE: Recursion is used by default, which resolves CNAME entries
		// back to the actual A record
		msg.SetQuestion(fqdn, dns.TypeA)

		// Perform UDP-based query using default settings
		in, err := dns.Exchange(&msg, server+":53")
		if err != nil {
			panic(err)
		}

		// Early exit if one of the DNS servers returns an unexpected result
		if len(in.Answer) < 1 {
			fmt.Printf("ERROR: No records for %q from %s\n", cfg.Query, server)
			return
		}

		dnsQueryResponse := DNSQueryResponse{
			// use zero value initially for Answer field
			Answer: in.Answer,
			Server: server,
			Query:  cfg.Query,
		}

		results = append(results, dnsQueryResponse)
	}

	// loop over all collected query responses

	w := new(tabwriter.Writer)
	//w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, '.', tabwriter.AlignRight|tabwriter.Debug)

	// Format in tab-separated columns
	//w.Init(os.Stdout, 16, 8, 8, '\t', 0)
	w.Init(os.Stdout, 8, 8, 8, '\t', 0)

	// Header row in output
	fmt.Fprintln(w,
		"Server\tQuery\tAnswers\tTTL\t")

	fmt.Fprintln(w,
		"---\t---\t---\t---\t")

	for _, item := range results {
		fmt.Fprintf(w,
			"%s\t%s\t%s\t%s\n",
			item.Server,
			item.Query,
			item.Records(),
			item.TTLs(),
		)
	}

	fmt.Fprintln(w)
	w.Flush()

}
