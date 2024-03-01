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

// PrintSummary generates a summary of all collected DNS query results in the
// specified format. If specified, the date/time that the results are
// generated is omitted from the results output.
func (dqrs DNSQueryResponses) PrintSummary(outputFormat string, omitTimestamp bool) {

	w := tabwriter.NewWriter(os.Stdout, 4, 4, 4, ' ', 0)

	// Add some lead-in spacing to better separate any earlier log messages from
	// summary output
	fmt.Fprintf(w, "\n\n")

	// REMINDER: Column cells must be tab-terminated, not tab-separated:
	// non-tab terminated trailing text at the end of a line forms a cell but
	// that cell is not part of an aligned column.
	var headerRowTmpl string
	var separatorRowTmpl string
	var recordRowErrorTmpl string
	var recordRowSuccessTmpl string

	switch {

	case !dqrs.RecordsFound():

		headerRowTmpl = "Server\tRTT\tQuery\tType\tAnswer\t"
		separatorRowTmpl = "---\t---\t---\t---\t---\t"
		recordRowErrorTmpl = "%s\t%s\t%s\t%s\t%s\t\n"
		recordRowSuccessTmpl = "%s\t%s\t%s\t%s\t%s\t\n"

	case outputFormat == ResultsOutputMultiLine:

		headerRowTmpl = "Server\tRTT\tQuery\tType\tAnswer\tAnswer Type\tTTL\t"
		separatorRowTmpl = "---\t---\t---\t---\t---\t---\t---\t"
		recordRowErrorTmpl = "%s\t%s\t%s\t%s\t%s\t\t\t\n"
		recordRowSuccessTmpl = "%s\t%s\t%s\t%s\t%s\t%s\t%d\t\n"

	case outputFormat == ResultsOutputSingleLine:
		headerRowTmpl = "Server\tRTT\tQuery\tType\tAnswers\tTTL\t"
		separatorRowTmpl = "---\t---\t---\t---\t---\t---\t"
		recordRowErrorTmpl = "%s\t%s\t%s\t%s\t%s\t\t\n"
		recordRowSuccessTmpl = "%s\t%s\t%s\t%s\t%s\t%s\t\n"

	}

	// Header row in output
	fmt.Fprintln(w, headerRowTmpl)

	// Separator row
	fmt.Fprintln(w, separatorRowTmpl)

	for _, item := range dqrs {
		// Building with `go build -gcflags=all=-d=loopvar=2` identified this
		// loop as compiling differently with Go 1.22 (per-iteration) loop
		// semantics. In particular, it is believed that the use of the Answer
		// ([]dns.RR) struct field which caused this loop to be flagged.
		//
		// As a workaround, we create a new variable for each iteration to
		// work around potential issues with Go versions prior to Go 1.22.
		item := item

		requestType, err := RRTypeToString(item.RequestedRecordType)
		if err != nil {
			requestType = "rrString LookupError"
		}

		// if any errors were recorded when querying DNS server show those
		// instead of attempting to show real results
		if item.QueryError != nil {
			fmt.Fprintf(w,
				recordRowErrorTmpl,
				item.Server,
				item.ResponseTime.Round(time.Millisecond),
				item.Query,
				requestType,
				item.QueryError.Error(),
			)
			continue
		}

		// Sort records before printing them
		item.SortRecordsAsc()

		switch outputFormat {
		case ResultsOutputMultiLine:

			for _, record := range item.Records() {
				fmt.Fprintf(w,
					recordRowSuccessTmpl,
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

		case ResultsOutputSingleLine:

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
				recordRowSuccessTmpl,
				item.Server,
				item.ResponseTime.Round(time.Millisecond),
				item.Query,
				requestType,
				strings.Join(responses, ", "),
				strings.Join(ttls, ", "),
			)
		}

	}

	if !omitTimestamp {
		fmt.Fprintf(
			w,
			"\nQuery Performed: %v",
			time.Now().Format(time.RFC3339),
		)
	}

	fmt.Fprintln(w)

	if err := w.Flush(); err != nil {
		log.Errorf("Error flushing tabwriter: %v", err.Error())
	}

}
