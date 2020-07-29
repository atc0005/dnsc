// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/nagios-debug
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {

	numEnvVars := len(os.Environ())

	fmt.Printf(
		"Total environment variables: %v",
		numEnvVars,
	)

	fmt.Printf("Environment variables:\r\n\r\n")

	origEnvVars := os.Environ()
	sortedEnvVars := make([]string, len(origEnvVars))
	copy(sortedEnvVars, origEnvVars)

	sort.Slice(sortedEnvVars, func(i, j int) bool {
		return sortedEnvVars[i] < sortedEnvVars[j]
	})

	for _, e := range sortedEnvVars {
		fmt.Println(e)
	}

}
