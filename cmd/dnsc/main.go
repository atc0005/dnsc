// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/dnsc
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"sort"

	"github.com/miekg/dns"

	"github.com/atc0005/dnsc/config"
)

// DNSQueryResponse represents a query and response from a DNS server
type DNSQueryResponse struct {

	// Depending on the system, there *could* be multiple A records returned
	ARecords []net.IP
	Server   string
	Query    string
}

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// do something useful with the config
	fmt.Printf("%#v\n", *cfg)

	// Process query using configured settings

	// fqdn := dns.Fqdn("stacktitan.com")
	fqdn := dns.Fqdn(cfg.Query)

	// loop over each of our DNS servers, build up a results set

	var results []DNSQueryResponse

	for _, server := range cfg.Servers {
		//fmt.Println("Server:", server)

		var msg dns.Msg

		// NOTE: Recursion is used by default, which resolves CNAME entries
		// back to the actual A record
		msg.SetQuestion(fqdn, dns.TypeA)

		in, err := dns.Exchange(&msg, server+":53")
		if err != nil {
			panic(err)
		}

		if len(in.Answer) < 1 {
			fmt.Printf("DEBUG: No records for %q from %s\n", cfg.Query, server)
			return
		}

		dnsQueryResponse := DNSQueryResponse{
			// use zero value initially for Answer field
			//Answer:,
			Server: server,
			Query:  cfg.Query,
		}

		//fmt.Println("length of in.Answer:", len(in.Answer))
		//fmt.Printf("%#v\n", in.Answer)

		// We could get back a CNAME entry in addition to an A record, so loop
		// until we find an A record
		typeARecordsFound := make([]net.IP, 0, 5)
		for _, answer := range in.Answer {
			if a, ok := answer.(*dns.A); ok {
				typeARecordsFound = append(typeARecordsFound, a.A)
			}
		}

		// Sort to make comparison easier later
		// https://stackoverflow.com/a/48389676
		sort.Slice(typeARecordsFound, func(i, j int) bool {
			return bytes.Compare(typeARecordsFound[i], typeARecordsFound[j]) < 0
		})

		dnsQueryResponse.ARecords = typeARecordsFound

		results = append(results, dnsQueryResponse)
	}

	// loop over all collected query responses

	for _, item := range results {
		//fmt.Printf("%#v\n", item)
		fmt.Printf(
			"Server: %s, Query: %s, Answer: %v\n",
			item.Server,
			item.Query,
			item.ARecords,
		)
	}

}
