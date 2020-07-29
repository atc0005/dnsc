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

	fmt.Printf("Environment variables (%d):\n\n", numEnvVars)

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
