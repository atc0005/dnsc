# dnsc

Queries a list of DNS servers and compares responses against an expected answer

## Project home

See [our GitHub repo](https://github.com/atc0005/dnsc) for the latest
code, to file an issue or submit improvements for review and potential
inclusion into the project.

## Overview

Run a DNS query against all servers in a list and compare against an expected
value. This check is most useful after moving servers between subnets when an
IP Address change is expected.

TODO:

- Note available command-line flags are for options which frequently change
- Config file for infrequently changing settings

## Configuration

## Examples

These examples assume that the referenced `config.toml` file has the following
contents:

```toml

[dns]

# Multi-line array
# https://github.com/toml-lang/toml#user-content-array
servers = [
    "8.8.8.8",
    "8.8.4.4",
]

```

```ShellSession
dnsc -config-file $HOME/.config/dnsc/config.toml -q server1.example.com -e 192.168.1.250
```

```ShellSession
dnsc -config-file $HOME/.config/dnsc/config.toml -query server1.example.com -expect 192.168.1.250
```

## Expected results

- Should print small table of details with a final pass/fail verdict
- TTL for each returned item
- Expected answer
- Actual answer

## References

- Golang basics
  - <https://golang.org/pkg/sort/#Slice>
  - <https://golang.org/pkg/flag/#example_>
  - <https://stackoverflow.com/questions/24886015/how-to-convert-uint32-to-string>
  - <https://www.geeksforgeeks.org/nested-structure-in-golang/>

- TOML config file-related
  - <https://github.com/toml-lang/toml#user-content-array>
  - <https://github.com/pelletier/go-toml>

- DNS library
  - <https://github.com/miekg/dns>
  - <https://miek.nl/2014/august/16/go-dns-package/>
  - <https://nostarch.com/blackhatgo>
    - <https://github.com/blackhat-go/bhg/blob/master/ch-5/get_all_a/main.go>

- Free DNS servers for testing purposes
  - <https://developers.google.com/speed/public-dns>
  - <https://www.opendns.com/setupguide/>
