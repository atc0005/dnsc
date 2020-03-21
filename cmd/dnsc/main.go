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
	"sort"

	"github.com/atc0005/dnsc/config"
	"github.com/atc0005/dnsc/dqrs"
	"github.com/miekg/dns"

	"github.com/apex/log"
)

func main() {

	// This will use default logging settings (level filter, destination)
	// as the application hasn't "booted up" far enough to apply custom
	// choices yet.
	log.Debug("Initializing application")

	cfg, err := config.NewConfig()
	switch {
	// TODO: How else to guard against nil cfg object?
	case cfg != nil && cfg.ShowVersion():
		config.Branding()
		os.Exit(0)
	case err == nil:
		// do nothing for this one
	case errors.Is(err, flag.ErrHelp):
		os.Exit(0)
	default:
		log.Fatalf("failed to initialize application: %s", err)
	}

	// TODO: Move this elsewhere, add support for flag and TOML setting
	requestTypes := []unint16{
		dns.TypeA,
		dns.TypeAAAA,
	}

	expectedResponses := len(requestTypes) * len(cfg.Servers())

	results := make(dqrs.DNSQueryResponses, 0, 10)

	capacity := len(cfg.Servers())
	log.WithFields(log.Fields{
		"results_channel_capacity": capacity,
	}).Debug("Creating results channel with capacity to match defined DNS servers")
	// resultsChan := make(chan dqrs.DNSQueryResponse)
	resultsChan := make(chan dqrs.DNSQueryResponse, capacity)

	// spin off a separate goroutine for each of our DNS servers, send back
	// results on a channel
	for _, server := range cfg.Servers() {

		go func(server string, query string, results chan dqrs.DNSQueryResponse) {
			// dnsQueryResponse := dqrs.PerformQuery(query, server, dns.TypeA)
			resultsChan <- dqrs.PerformQuery(query, server, dns.TypeA)
			resultsChan <- dqrs.PerformQuery(query, server, dns.TypeAAAA)
		}(server, cfg.Query(), resultsChan)

	}

	// Collect all responses using the total number of DNS servers as limiter
	// multiplied by the number of questions we have for each DNS server
	remainingResponses := len(cfg.Servers()) * 2
	for remainingResponses > 0 {
		result := <-resultsChan
		results = append(results, result)
		if result.QueryError != nil {
			// Check whether the user has opted to ignore errors. If not,
			// display current summary results and exit
			if !cfg.IgnoreDNSErrors() {
				results.PrintSummary()
				os.Exit(1)
			}
		}

		// note that we've received another response from a DNS server in our
		// list; get the next response, break out of loop once done
		remainingResponses--
	}

	// Sort DNS query results results by server used for query. This is done
	// in an effort to arrange responses based on the group of DNS servers
	// (assuming that they're group together using a consecutive IP block)
	sort.Slice(results, func(i, j int) bool {
		return results[i].Server < results[j].Server
	})

	// Generate summary of all collected query responses
	results.PrintSummary()

}
