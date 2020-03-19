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

### Precedence

**Note: This behavior is subject to change based on feedback.**

The priority order is (mostly):

1. Command line flags (highest priority)
1. Configuration file
1. Default settings (lowest priority)

The intent of this behavior is to provide a *feathered* layering of
configuration settings; if a configuration file provides all settings that you
want other than one, you can use the configuration file for the other settings
and specify the settings that you wish to override via command-line flag.

### Command-line arguments

Notes:

- As of this writing, a config-file *is* required to specify DNS servers. The
  plan is to add support for specifying this via flag, but it hasn't been
  added yet.

- Flags marked as required *are* required to be set via CLI flag *or* within a
  TOML-formatted configuration file.

| Flag                       | Required | Default        | Repeat | Possible                                   | Description                                                                                                              |
| -------------------------- | -------- | -------------- | ------ | ------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `h`, `help`                | No       | `false`        | No     | `h`, `help`                                | Show Help text along with the list of supported flags.                                                                   |
| `cf`, `config-file`        | **Yes**  | *empty string* | No     | *valid file name characters*               | Full path to TOML-formatted configuration file. See [`config.example.toml`](config.example.toml) for a starter template. |
| `v`, `version`             | No       | `false`        | No     | `true`, `false`                            | Whether to display application version and then immediately exit application.                                            |
| `ide`, `ignore-dns-errors` | No       | `false`        | No     | `true`, `false`                            | Whether DNS-related errors with one server should be ignored in order to try other DNS servers in the list.              |
| `q`, `query`               | **Yes**  | *empty string* | No     | *any valid FQDN string*                    | Fully-qualified system to lookup from all provided DNS servers.                                                          |
| `ll`, `log-level`          | No       | `info`         | No     | `fatal`, `error`, `warn`, `info`, `debug`  | Log message priority filter. Log messages with a lower level are ignored.                                                |
| `lf`, `log-format`         | No       | `text`         | No     | `cli`, `json`, `logfmt`, `text`, `discard` | Use the specified `apex/log` package "handler" to output log messages in that handler's format.                          |


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
