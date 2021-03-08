// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/dnsc
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package dqrs

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/apex/log"
)

func (dqrs DNSQueryResponses) resultsSummaryOutputMultiLine() {

	w := tabwriter.NewWriter(os.Stdout, 4, 4, 4, ' ', 0)

	// Add some lead-in spacing to better separate any earlier log messages from
	// summary output
	fmt.Fprintf(w, "\n\n")

	// Header row in output
	fmt.Fprintf(w,
		"Server\tRTT\tQuery\tQuery Type\tAnswer\tAnswer Type\tTTL\t\n")

	// Separator row
	fmt.Fprintln(w,
		"---\t---\t---\t---\t---\t---\t---\t")

	for _, item := range dqrs {

		requestType, err := RRTypeToString(item.RequestedRecordType)
		if err != nil {
			requestType = "rrString LookupError"
		}

		// if any errors were recorded when querying DNS server show those
		// instead of attempting to show real results
		if item.QueryError != nil {
			fmt.Fprintf(w,
				"%s\t%s\t%s\t%s\t%s\t%s\t%s\t\n",
				item.Server,
				item.ResponseTime.Round(time.Millisecond),
				item.Query,
				requestType,
				item.QueryError.Error(),
				"",
				"",
			)
			continue
		}

		// Sort records before printing them
		item.SortRecordsAsc()

		for _, record := range item.Records() {
			fmt.Fprintf(w,
				"%s\t%s\t%s\t%s\t%s\t%s\t%d\t\n",
				item.Server,
				item.ResponseTime.Round(time.Millisecond),
				item.Query,
				requestType,
				record.Value,

				// Display request type from record, which may not match the
				// original request type (e.g., CNAME and A records returned
				// for original lookup.
				record.Type,
				record.TTL,
			)
		}

	}

	fmt.Fprintln(w)

	if err := w.Flush(); err != nil {
		log.Errorf("Error flushing tabwriter: %v", err.Error())
	}
}

func (dqrs DNSQueryResponses) resultsSummaryOutputSingleLine() {

	w := tabwriter.NewWriter(os.Stdout, 4, 4, 4, ' ', 0)

	// Add some lead-in spacing to better separate any earlier log messages from
	// summary output
	fmt.Fprintf(w, "\n\n")

	// Header row in output
	fmt.Fprintf(w,
		"Server\tRTT\tQuery\tType\tAnswers\tTTL\t\n")

	// Separator row
	fmt.Fprintln(w,
		"---\t---\t---\t---\t---\t---\t")

	for _, item := range dqrs {

		rrString, err := RRTypeToString(item.RequestedRecordType)
		if err != nil {
			rrString = "rrString LookupError"
		}

		// if any errors were recorded when querying DNS server show those
		// instead of attempting to show real results
		if item.QueryError != nil {
			fmt.Fprintf(w,
				"%s\t%s\t%s\t%s\t%s\t%s\t\n",
				item.Server,
				item.ResponseTime.Round(time.Millisecond),
				item.Query,
				rrString,
				item.QueryError.Error(),
				"",
			)
			continue
		}

		// Sort records before printing them
		item.SortRecordsAsc()

		var responses []string
		var ttls []string
		for _, record := range item.Records() {
			response := fmt.Sprintf(
				"%s (%s)",
				record.Value,
				record.Type,
			)
			responses = append(responses, response)
			ttls = append(ttls, fmt.Sprint(record.TTL))
		}

		fmt.Fprintf(w,
			"%s\t%s\t%s\t%s\t%s\t%s\t\n",
			item.Server,
			item.ResponseTime.Round(time.Millisecond),
			item.Query,
			rrString,
			strings.Join(responses, ", "),
			strings.Join(ttls, ", "),
		)
	}

	fmt.Fprintln(w)

	if err := w.Flush(); err != nil {
		log.Errorf("Error flushing tabwriter: %v", err.Error())
	}
}

// PrintSummary generates a summary of all collected DNS query results in the
// specified format.
func (dqrs DNSQueryResponses) PrintSummary(outputFormat string) {

	switch outputFormat {

	case ResultsOutputMultiLine:
		dqrs.resultsSummaryOutputMultiLine()
	case ResultsOutputSingleLine:
		dqrs.resultsSummaryOutputSingleLine()
	default:
		log.Warnf("Unknown results output format specified: %s", outputFormat)
		log.Warnf("Defaulting to %s output", ResultsOutputSingleLine)
		dqrs.resultsSummaryOutputSingleLine()
	}

}
