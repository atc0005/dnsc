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

	"github.com/miekg/dns"

	"github.com/atc0005/dnsc/config"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fqdn := dns.Fqdn(cfg.Query)

	results := make(DNSQueryResponses, 0, 10)

	// loop over each of our DNS servers, build up a results set
	for _, server := range cfg.Servers {

		var msg dns.Msg

		// NOTE: Recursion is used by default, which resolves CNAME entries
		// back to the actual A record
		msg.SetQuestion(fqdn, dns.TypeA)

		// TODO: Use concurrency to query all DNS servers in bulk instead of
		// waiting for each to complete before querying the next one

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

	// Generate summary of all collected query responses
	results.PrintSummary()

}
