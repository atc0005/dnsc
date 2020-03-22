/*

Submit query against a list of DNS servers and display summary of results

PROJECT HOME

See our GitHub repo (https://github.com/atc0005/dnsc) for the latest
code, to file an issue or submit improvements for review and potential
inclusion into the project.

PURPOSE

Run a DNS query concurrently against all servers in a list and provide summary
of results. This is most useful after moving servers between subnets when an
IP Address change is expected, but where the change may not have propagated
between all DNS servers. The summary output is useful for spotting systems
lagging behind the others.

Command-line flags are supported for all options, though for some settings
(e.g., DNS servers), specifying values via configuration file is easier for
repeat use.

FEATURES

• single binary, no outside dependencies

• User configurable logging levels

• User configurable logging format

USAGE

Help output is below. See the README for examples.

$ ./dnsc.exe -h

    dnsc dev build
    https://github.com/atc0005/dnsc

    Usage of "dnsc.exe":
    -cf string
        Full path to TOML-formatted configuration file. See config.example.toml for a starter template. (shorthand)
    -config-file string
        Full path to TOML-formatted configuration file. See config.example.toml for a starter template.
    -dns-server value
        DNS server to submit query against. This flag may be repeated for each additional DNS server to query.
    -ds value
        DNS server to submit query against. This flag may be repeated for each additional DNS server to query. (shorthand)
    -ide
        Whether DNS-related errors with one server should be ignored in order to try other DNS servers in the list. (shorthand)
    -ignore-dns-errors
        Whether DNS-related errors with one server should be ignored in order to try other DNS servers in the list.
    -lf string
        Log messages are written in this format (default "text")
    -ll string
        Log message priority filter. Log messages with a lower level are ignored. (default "info")
    -log-format string
        Log messages are written in this format (default "text")
    -log-level string
        Log message priority filter. Log messages with a lower level are ignored. (default "info")
    -q string
        Fully-qualified system to lookup from all provided DNS servers. (shorthand)
    -query string
            Fully-qualified system to lookup from all provided DNS servers.
    -t value
        DNS query type to use when submitting DNS queries. The default is the 'A' query type. This flag may be repeated for each additional DNS record type you wish to request. (shorthand)
    -type value
        DNS query type to use when submitting DNS queries. The default is the 'A' query type. This flag may be repeated for each additional DNS record type you wish to request.
    -v
        Whether to display application version and then immediately exit application. (shorthand)
    -version
        Whether to display application version and then immediately exit application.

*/
package main
