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
	"sync"

	"github.com/atc0005/dnsc/config"
	"github.com/atc0005/dnsc/dqrs"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	var wg sync.WaitGroup

	// Declare that we'll have the same number of goroutines as we do the
	// number of DNS servers
	wg.Add(len(cfg.Servers))

	results := make(dqrs.DNSQueryResponses, 0, 10)

	// loop over each of our DNS servers, build up a results set
	for _, server := range cfg.Servers {

		go func(server string, query string, wg *sync.WaitGroup) {

			defer wg.Done()

			dnsQueryResponse, err := dqrs.PerformQuery(query, server)
			if err != nil {
				// Check whether the user has opted to ignore errors and proceed
				// FIXME: Assuming 'Yes' for now
				log.Println(err)
			}

			results = append(results, dnsQueryResponse)

		}(server, cfg.Query, &wg)

	}

	wg.Wait()

	// Generate summary of all collected query responses
	results.PrintSummary()

}
