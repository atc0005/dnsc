// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/dnsc
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"log"
	"os"

	"github.com/atc0005/dnsc/config"
	"github.com/atc0005/dnsc/dqrs"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	results := make(dqrs.DNSQueryResponses, 0, 10)

	// receive query results on this channel
	resultsChan := make(chan dqrs.DNSQueryResponse)

	// loop over each of our DNS servers, build up a results set
	for _, server := range cfg.Servers {

		go func(server string, query string, results chan dqrs.DNSQueryResponse) {

			dnsQueryResponse, err := dqrs.PerformQuery(query, server)
			if err != nil {
				// Check whether the user has opted to ignore errors and proceed
				// FIXME: Assuming 'Yes' for now
				log.Println(err)
			}

			//results = append(results, dnsQueryResponse)
			resultsChan <- dnsQueryResponse

		}(server, cfg.Query, resultsChan)

		// TODO: Signal that we're done spinning off goroutines?

	}

	// loop over results channel and collect all responses
	for response := range resultsChan {
		results = append(results, response)
	}

	// Generate summary of all collected query responses
	results.PrintSummary()

}
