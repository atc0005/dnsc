/*
Submit query against a list of DNS servers and display summary of results

# Project Home

See our GitHub repo (https://github.com/atc0005/dnsc) for the latest
code, to file an issue or submit improvements for review and potential
inclusion into the project.

# Purpose

Run a DNS query concurrently against all servers in a list and provide summary
of results. This is most useful after moving servers between subnets when an
IP Address change is expected, but where the change may not have propagated
between all DNS servers. The summary output is useful for spotting systems
lagging behind the others.

Command-line flags are supported for all options, though for some settings
(e.g., DNS servers), specifying values via configuration file is easier for
repeat use.

# Features

  - single binary, no outside dependencies
  - Multiple query types supported
  - (Optional) SRV record protocol "shortcuts" for more complex queries
  - User configurable logging levels
  - User configurable logging format
  - User configurable results summary output format
  - User configurable query timeout

# Usage

See the README for examples.
*/
package main
