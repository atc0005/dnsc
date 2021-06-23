// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/dnsc
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"sync"
	"time"

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

	cfg, cfgErr := config.NewConfig()
	switch {
	case errors.Is(cfgErr, config.ErrVersionRequested):
		config.Branding()
		os.Exit(0)
	case cfgErr != nil:
		log.Fatalf("failed to initialize application: %s", cfgErr)
	}

	// Get a list of all record types that we should request when submitting
	// DNS queries
	queryTypes := cfg.QueryTypes()
	queryTimeout := cfg.Timeout()

	var expectedResponses int
	switch {
	case len(cfg.SrvProtocols()) > 0:
		types := len(queryTypes) - 1         // exclude SRV record query type
		protocols := len(cfg.SrvProtocols()) // SRV query type calculated here
		servers := len(cfg.Servers())

		expectedResponses = (types * servers) + (protocols * servers)
	default:
		expectedResponses = len(queryTypes) * len(cfg.Servers())
	}

	log.Debugf("%d queries to submit, equal number responses expected\n", expectedResponses)

	results := make(dqrs.DNSQueryResponses, 0, expectedResponses)
	resultsChan := make(chan dqrs.DNSQueryResponse)

	var queriesWG sync.WaitGroup
	var collectorWG sync.WaitGroup

	// spin off "collector" for results channel
	collectorWG.Add(1)
	go func() {
		defer collectorWG.Done()
		// Collect all responses
		for result := range resultsChan {

			log.Debug("collector: Received result")

			results = append(results, result)
			if result.QueryError != nil {
				// Check whether the user has opted to treat errors as fatal. If
				// so, display current summary results and exit
				if cfg.DNSErrorsFatal() {
					results.PrintSummary(cfg.ResultsOutput())
					os.Exit(1)
				}
			}

			log.Debug("collector: Saved result")
		}

		log.Debug("collector: Exited collection loop")
	}()

	// spin off a separate goroutine for each of our DNS servers, send back
	// results on a channel
	for _, server := range cfg.Servers() {

		queriesWG.Add(1)
		go func(server string, query string, queryTypes []string, queryTimeout time.Duration, results chan dqrs.DNSQueryResponse) {

			defer queriesWG.Done()

			log.Debugf("Length of requested query types: %d", len(queryTypes))
			for _, rrString := range queryTypes {
				rrType, err := dqrs.RRStringToType(rrString)
				if err != nil {
					// Record the error, log the error and send a minimal
					// DNSQueryResponse type back on the channel with the
					// specific error embedded.
					failedQueryRequest := dqrs.DNSQueryResponse{
						Server:     server,
						Query:      query,
						QueryError: fmt.Errorf("error converting Resource Record string to native type: %w", err),
					}
					log.Warn(failedQueryRequest.Error())
					queriesWG.Add(1)
					go func() {
						defer queriesWG.Done()
						resultsChan <- failedQueryRequest
					}()

				}

				var totalQueries int
				switch {
				case rrType == dns.TypeSRV && len(cfg.SrvProtocols()) > 0:
					totalQueries = len(cfg.SrvProtocols())
				default:
					totalQueries = 1
				}
				queries := make([]string, 0, totalQueries)

				switch {

				// If performing SRV record queries, check to see if SRV
				// protocols were specified. If so, submit a separate query
				// for each one after resolving the protocol record syntax
				// needed.
				case rrType == dns.TypeSRV && len(cfg.SrvProtocols()) > 0:
					for _, srvProtocol := range cfg.SrvProtocols() {
						queryTemplate, err := config.SrvProtocolTmplLookup(srvProtocol)
						if err != nil {
							// Record the error, log the error and send a
							// minimal DNSQueryResponse type back on the
							// channel with the specific error embedded.
							failedQueryRequest := dqrs.DNSQueryResponse{
								Server:     server,
								Query:      query,
								QueryError: fmt.Errorf("error retrieving SRV protocol template: %w", err),
							}
							log.Warn(failedQueryRequest.Error())
							queriesWG.Add(1)
							go func() {
								defer queriesWG.Done()
								resultsChan <- failedQueryRequest
							}()

						}
						queries = append(queries, fmt.Sprintf(queryTemplate, query))
					}
				// use the query string as-is if SRV protocols not specified
				default:
					queries = append(queries, query)
				}

				log.Debugf("Total queries collected: %d", len(queries))

				for i := range queries {
					log.Debugf("Submitting query for %q of type %q to %q",
						queries[i], rrString, server)

					queriesWG.Add(1)
					go func(q string) {
						defer queriesWG.Done()
						resultsChan <- dqrs.PerformQuery(q, server, rrType, queryTimeout)
						log.Debug("Query completed, results sent back on channel")
					}(queries[i])

				}

			}
		}(server, cfg.Query(), queryTypes, queryTimeout, resultsChan)

	}

	// Close results channel after all queries have been submitted and results
	// are ready to be collected
	queriesWG.Wait()
	close(resultsChan)

	// Wait on collector to finish saving results from earlier queries
	collectorWG.Wait()

	// Sort DNS query results by server used for query. This is done in an
	// effort to arrange responses based on the group of DNS servers (assuming
	// that they're grouped together using a consecutive IP block)
	sort.Slice(results, func(i, j int) bool {
		return results[i].Server < results[j].Server
	})

	// Generate summary of all collected query responses in the specified
	// format
	results.PrintSummary(cfg.ResultsOutput())

}
