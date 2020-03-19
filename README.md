# dnsc

Submit query against a list of DNS servers and display summary of results

- [dnsc](#dnsc)
  - [Project home](#project-home)
  - [Overview](#overview)
  - [Configuration](#configuration)
    - [Precedence](#precedence)
    - [Command-line arguments](#command-line-arguments)
    - [Configuration file](#configuration-file)
  - [Examples](#examples)
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

| Flag                       | Required | Default        | Repeat  | Possible                                   | Description                                                                                                              |
| -------------------------- | -------- | -------------- | ------- | ------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `h`, `help`                | No       | `false`        | No      | `h`, `help`                                | Show Help text along with the list of supported flags.                                                                   |
| `ds`, `dns-server`         | **Yes**  | *empty string* | **Yes** | *one valid IP Address per flag invocation* | DNS server to submit query against. This flag may be repeated for each additional DNS server to query.                   |
| `cf`, `config-file`        | **Yes**  | *empty string* | No      | *valid file name characters*               | Full path to TOML-formatted configuration file. See [`config.example.toml`](config.example.toml) for a starter template. |
| `v`, `version`             | No       | `false`        | No      | `v`, `version`                             | Whether to display application version and then immediately exit application.                                            |
| `ide`, `ignore-dns-errors` | No       | `false`        | No      | `ide`, `ignore-dns-errors`                 | Whether DNS-related errors with one server should be ignored in order to try other DNS servers in the list.              |
| `q`, `query`               | **Yes**  | *empty string* | No      | *any valid FQDN string*                    | Fully-qualified system to lookup from all provided DNS servers.                                                          |
| `ll`, `log-level`          | No       | `info`         | No      | `fatal`, `error`, `warn`, `info`, `debug`  | Log message priority filter. Log messages with a lower level are ignored.                                                |
| `lf`, `log-format`         | No       | `text`         | No      | `cli`, `json`, `logfmt`, `text`, `discard` | Use the specified `apex/log` package "handler" to output log messages in that handler's format.                          |

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

See the [`config.example.toml`](config.example.toml) file for an example of
how to use these settings.

## Examples

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

```ShellSession
$ dnsc -config-file $HOME/.config/dnsc/config.toml -q www.yahoo.com

    Server            Query            Answers                                                                                                            TTL
    ---               ---              ---                                                                                                                ---
    208.67.220.220    www.yahoo.com    atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 98.138.219.231 (A), 98.138.219.232 (A), 72.30.35.9 (A), 72.30.35.10 (A)    62, 30, 30, 30, 30
    208.67.222.222    www.yahoo.com    atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 72.30.35.9 (A), 72.30.35.10 (A), 98.138.219.231 (A), 98.138.219.232 (A)    1560, 40, 40, 40, 40
    8.8.4.4           www.yahoo.com    atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 72.30.35.9 (A), 72.30.35.10 (A)                                            515, 44, 44
    8.8.8.8           www.yahoo.com    atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 72.30.35.9 (A), 72.30.35.10 (A)                                            872, 17, 17
```

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
