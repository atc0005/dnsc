// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/dnsc
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	//"log"
	"errors"
	"flag"
	"os"

	"github.com/atc0005/dnsc/config"
	"github.com/atc0005/dnsc/dqrs"

	"github.com/apex/log"
	// "github.com/apex/log/handlers/cli"
	// "github.com/apex/log/handlers/discard"
	// "github.com/apex/log/handlers/json"
	// "github.com/apex/log/handlers/logfmt"
	// "github.com/apex/log/handlers/text"
)

func main() {

	// This will use default logging settings (level filter, destination)
	// as the application hasn't "booted up" far enough to apply custom
	// choices yet.
	log.Debug("Initializing application")

	cfg, err := config.NewConfig()
	switch {
	case cfg.ShowVersion:
		config.Branding()
		os.Exit(0)
	case err == nil:
		// do nothing for this one
	case errors.Is(err, flag.ErrHelp):
		os.Exit(0)
	default:
		log.Fatalf("Failed to initialize application: %s", err)
	}

	results := make(dqrs.DNSQueryResponses, 0, 10)

	// receive query results on this channel
	resultsChan := make(chan dqrs.DNSQueryResponse)

	// loop over each of our DNS servers, build up a results set
	for _, server := range cfg.Servers {

		go func(server string, query string, results chan dqrs.DNSQueryResponse) {

			dnsQueryResponse := dqrs.PerformQuery(query, server)
			if dnsQueryResponse.QueryError != nil {
				// Check whether the user has opted to ignore errors and proceed
				// FIXME: Assuming 'No' for now
				log.Error(dnsQueryResponse.QueryError.Error())
				os.Exit(1)
			}

			//results = append(results, dnsQueryResponse)
			resultsChan <- dnsQueryResponse

		}(server, cfg.Query, resultsChan)

		// TODO: Signal that we're done spinning off goroutines?

	}

	// collect all responses using the total number of DNS servers as our
	// limiter (for now, until I learn more about channels)
	remainingResponses := len(cfg.Servers)
	for remainingResponses > 0 {
		results = append(results, <-resultsChan)
		remainingResponses--
	}
	// for response := range resultsChan {
	// 	results = append(results, response)
	// }

	// Generate summary of all collected query responses
	results.PrintSummary()

}
