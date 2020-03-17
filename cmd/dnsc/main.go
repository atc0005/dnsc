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
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/miekg/dns"

	"github.com/atc0005/dnsc/config"
)

// DNSQueryResponse represents a query and response from a DNS server
type DNSQueryResponse struct {

	// Depending on the system, there *could* be multiple A records returned
	ARecords []*dns.A
	Server   string
	Query    string
}

func (dqr DNSQueryResponse) GetARecords() string {

	var aRecords []string
	for _, record := range dqr.ARecords {
		aRecords = append(aRecords, record.A.String())
	}

	return strings.Join(aRecords, ",")
}

func (dqr DNSQueryResponse) GetTTLs() string {

	// https://stackoverflow.com/questions/24886015/how-to-convert-uint32-to-string
	var ttlEntries []string
	for _, record := range dqr.ARecords {
		ttlEntries = append(ttlEntries, fmt.Sprint(record.Hdr.Ttl))
	}

	return strings.Join(ttlEntries, ",")
}

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// do something useful with the config
	//fmt.Printf("%#v\n", *cfg)

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
			fmt.Printf("ERROR: No records for %q from %s\n", cfg.Query, server)
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
		typeARecordsFound := make([]*dns.A, 0, 5)
		for _, answer := range in.Answer {
			if a, ok := answer.(*dns.A); ok {
				// fmt.Println("TTL:", a.Hdr.Ttl)
				// fmt.Println("A:", a.A)
				typeARecordsFound = append(typeARecordsFound, a)
			}
		}

		// Sort to make comparison easier later
		// https://stackoverflow.com/a/48389676
		sort.Slice(typeARecordsFound, func(i, j int) bool {
			return bytes.Compare(typeARecordsFound[i].A, typeARecordsFound[j].A) < 0
		})

		dnsQueryResponse.ARecords = typeARecordsFound

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
		"Server\tQuery\tAnswers\tExpected\tTTL\t")

	for _, item := range results {
		fmt.Fprintf(w,
			"%s\t%s\t%s\t%s\t%s\n",
			item.Server,
			item.Query,
			item.GetARecords(),
			cfg.ExpectedAnswer,
			item.GetTTLs(),
		)
	}

	fmt.Fprintln(w)
	w.Flush()

}
