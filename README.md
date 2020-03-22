# dnsc

Submit query against a list of DNS servers and display summary of results

- [dnsc](#dnsc)
  - [Project home](#project-home)
  - [Overview](#overview)
  - [Features](#features)
    - [Current](#current)
    - [Planned](#planned)
  - [Configuration](#configuration)
    - [Precedence](#precedence)
    - [Command-line arguments](#command-line-arguments)
    - [Configuration file](#configuration-file)
  - [Examples](#examples)
    - [Setup config file](#setup-config-file)
    - [Use config file for DNS servers list](#use-config-file-for-dns-servers-list)
    - [Specify DNS servers list via flags](#specify-dns-servers-list-via-flags)
    - [Ignore errors, query all servers](#ignore-errors-query-all-servers)
  - [References](#references)

## Project home

See [our GitHub repo](https://github.com/atc0005/dnsc) for the latest
code, to file an issue or submit improvements for review and potential
inclusion into the project.

## Overview

Run a DNS query concurrently against all servers in a list and provide summary
of results. This is most useful after moving servers between subnets when an
IP Address change is expected, but where the change may not have propagated
between all DNS servers. The summary output is useful for spotting systems
lagging behind the others.

Command-line flags are supported for all options, though for some settings
(e.g., DNS servers), specifying values via configuration file is easier for
repeat use.

## Features

### Current

- A mix of command-line flags and configuration file options may be used for
  all options
- Query just one server or as many as are provided
  - Note: A configuration file is recommended for providing multiple DNS
    servers
- Multiple query types supported
  - `CNAME`
  - `A`
  - `AAAA`
  - `MX`

- User configurable logging levels

- User configurable logging format

### Planned

See [our GitHub repo](https://github.com/atc0005/dnsc) for planned future
work.

## Configuration

### Precedence

The priority order is:

1. Command line flags (highest priority)
1. Configuration file
1. Default settings (lowest priority)

The intent of this behavior is to provide a *feathered* layering of
configuration settings; if a configuration file provides nearly all settings
that you want, specify just the settings that you wish to override via
command-line flags and use the configuration file for the other settings.

### Command-line arguments

- Flags marked as **`required`** must be set via CLI flag *or* within a
  TOML-formatted configuration file.
- Flags *not* marked as required are for settings where a useful default is
  already defined.

| Flag                       | Required | Default        | Repeat  | Possible                                   | Description                                                                                                                                                   |
| -------------------------- | -------- | -------------- | ------- | ------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `h`, `help`                | No       | `false`        | No      | `h`, `help`                                | Show Help text along with the list of supported flags.                                                                                                        |
| `ds`, `dns-server`         | **Yes**  | *empty string* | **Yes** | *one valid IP Address per flag invocation* | DNS server to submit query against. This flag may be repeated for each additional DNS server to query.                                                        |
| `cf`, `config-file`        | **Yes**  | *empty string* | No      | *valid file name characters*               | Full path to TOML-formatted configuration file. See [`config.example.toml`](config.example.toml) for a starter template.                                      |
| `v`, `version`             | No       | `false`        | No      | `v`, `version`                             | Whether to display application version and then immediately exit application.                                                                                 |
| `ide`, `ignore-dns-errors` | No       | `false`        | No      | `ide`, `ignore-dns-errors`                 | Whether DNS-related errors with one server should be ignored in order to try other DNS servers in the list.                                                   |
| `q`, `query`               | **Yes**  | *empty string* | No      | *any valid FQDN string*                    | Fully-qualified system to lookup from all provided DNS servers.                                                                                               |
| `ll`, `log-level`          | No       | `info`         | No      | `fatal`, `error`, `warn`, `info`, `debug`  | Log message priority filter. Log messages with a lower level are ignored.                                                                                     |
| `lf`, `log-format`         | No       | `text`         | No      | `cli`, `json`, `logfmt`, `text`, `discard` | Use the specified `apex/log` package "handler" to output log messages in that handler's format.                                                               |
| `t`, `type`                | No       | `A`            | **Yes** | `A`, `AAAA`, `MX`, `CNAME`                 | DNS query type to use when submitting a DNS query to each provided server. This flag may be repeated for each additional DNS record type you wish to request. |

### Configuration file

Configuration file settings have the lowest priority and are overridden by
settings specified in other configuration sources, except for default values.
See the [Command-line Arguments](#command-line-arguments) table for more
information, including the available values for the listed configuration
settings.

| Flag Name           | Config file Setting Name | Notes                                                                              |
| ------------------- | ------------------------ | ---------------------------------------------------------------------------------- |
| `dns-server`        | `dns_servers`            | [Multi-line array](https://github.com/toml-lang/toml#user-content-array)           |
| `query`             | `query`                  | While supported, having a fixed query in the config file is not a normal use case. |
| `ignore-dns-errors` | `ignore_dns_errors`      | Opt-in setting. Useful to have enabled for most use cases.                         |
| `log-level`         | `log_level`              |                                                                                    |
| `log-format`        | `log_format`             |                                                                                    |
| `type`              | `dns_request_types`      | [Multi-line array](https://github.com/toml-lang/toml#user-content-array)           |

See the [`config.example.toml`](config.example.toml) file for an example of
how to use these settings.

## Examples

### Setup config file

These examples assume that the referenced `config.toml` file has the following
contents:

```toml

dns_servers = [
    # https://developers.google.com/speed/public-dns
    "8.8.8.8",
    "8.8.4.4",

    # https://www.opendns.com/setupguide/
    "208.67.222.222",
    "208.67.220.220",
]
```

See the [`config.example.toml`](config.example.toml) file for a starting point.

### Use config file for DNS servers list

```ShellSession
$ dnsc -config-file $HOME/.config/dnsc/config.toml -q www.yahoo.com

    Server            Query            Answers                                                                                                            TTL
    ---               ---              ---                                                                                                                ---
    208.67.220.220    www.yahoo.com    atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 98.138.219.231 (A), 98.138.219.232 (A), 72.30.35.9 (A), 72.30.35.10 (A)    62, 30, 30, 30, 30
    208.67.222.222    www.yahoo.com    atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 72.30.35.9 (A), 72.30.35.10 (A), 98.138.219.231 (A), 98.138.219.232 (A)    1560, 40, 40, 40, 40
    8.8.4.4           www.yahoo.com    atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 72.30.35.9 (A), 72.30.35.10 (A)                                            515, 44, 44
    8.8.8.8           www.yahoo.com    atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 72.30.35.9 (A), 72.30.35.10 (A)                                            872, 17, 17
```

### Specify DNS servers list via flags

You can also specify the DNS servers via CLI flags, though it is a bit more verbose:

```ShellSession
$ dnsc -ds 8.8.8.8 -ds 8.8.4.4 -ds 208.67.220.220 -ds 208.67.222.222 -q www.yahoo.com

    Server            Query            Answers                                                                                                            TTL
    ---               ---              ---                                                                                                                ---
    208.67.220.220    www.yahoo.com    atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 72.30.35.10 (A), 98.138.219.231 (A), 98.138.219.232 (A), 72.30.35.9 (A)    140, 53, 53, 53, 53
    208.67.222.222    www.yahoo.com    atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 72.30.35.9 (A), 72.30.35.10 (A), 98.138.219.231 (A), 98.138.219.232 (A)    464, 49, 49, 49, 49
    8.8.4.4           www.yahoo.com    atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 72.30.35.10 (A), 72.30.35.9 (A), 98.138.219.232 (A), 98.138.219.231 (A)    348, 54, 54, 54, 54
    8.8.8.8           www.yahoo.com    atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 72.30.35.10 (A), 72.30.35.9 (A), 98.138.219.232 (A), 98.138.219.231 (A)    347, 53, 53, 53, 53
```

### Ignore errors, query all servers

By default, one failure causes the application to immediately fail and query
results against other servers to be discarded. You can override this behavior
to allow the results to be processed from other servers. This can be useful if
you have one DNS server from the group unreachable or providing invalid
results.

Here we use the short-hand `-ide` flag:

```ShellSession
$ dnsc -ds 8.8.8.8 -ds 8.8.4.4 -ds 208.67.220.220 -ds 208.67.222.222 -q tacos -ide

    Server            Query    Answers                       TTL
    ---               ---      ---                           ---
    208.67.220.220    tacos    no records found for query
    208.67.222.222    tacos    no records found for query
    8.8.4.4           tacos    no records found for query
    8.8.8.8           tacos    no records found for query
```

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
