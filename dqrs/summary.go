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

	"github.com/miekg/dns"
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
		"Server\tQuery\tType\tAnswers\tTTL\n")

	// Separator row
	// TODO: I'm sure this can be handled better
	fmt.Fprintln(w,
		"---\t---\t---\t---\t---")

	for _, item := range dqrs {

		recordString, ok := dns.TypeToString[item.RequestedRecordType]
		if !ok {
			recordString = "unknown record type"
		}

		// if any errors were recorded when querying DNS server show those
		// instead of attempting to show real results
		if item.QueryError != nil {
			fmt.Fprintf(w,
				"%s\t%s\t%s\t%s\t%s\n",
				item.Server,
				item.Query,
				recordString,
				item.QueryError.Error(),
				"",
			)
			continue
		}

		// Sort records before printing them
		item.SortRecordsAsc()

		fmt.Fprintf(w,
			"%s\t%s\t%s\t%s\t%s\n",
			item.Server,
			item.Query,
			recordString,
			item.Records(),
			item.TTLs(),
		)
	}

	fmt.Fprintln(w)
	w.Flush()
}
