// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/dnsc
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package dqrs

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/apex/log"
)

// PrintSummary generates a table of all collected DNS query results
func (dqrs DNSQueryResponses) PrintSummary() {
	w := new(tabwriter.Writer)
	//w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, '.', tabwriter.AlignRight|tabwriter.Debug)

	// Format in tab-separated columns
	//w.Init(os.Stdout, 16, 8, 8, '\t', 0)
	w.Init(os.Stdout, 4, 4, 4, ' ', 0)

	// Add some lead-in spacing to better separate any earlier log messages from
	// summary output
	fmt.Fprintf(w, "\n\n")

	// Header row in output
	fmt.Fprintf(w,
		"Server\tRTT\tQuery\tType\tAnswers\tTTL\t\n")

	// Separator row
	// TODO: I'm sure this can be handled better
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

		fmt.Fprintf(w,
			"%s\t%s\t%s\t%s\t%s\t%s\t\n",
			item.Server,
			item.ResponseTime.Round(time.Millisecond),
			item.Query,
			rrString,
			item.Records(),
			item.TTLs(),
		)
	}

	fmt.Fprintln(w)

	// TODO: Add a retry?
	if err := w.Flush(); err != nil {
		log.Errorf("Error flushing tabwriter: %v", err.Error())
	}
}
