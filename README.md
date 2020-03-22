# dnsc

Submit query against a list of DNS servers and display summary of results

[![Latest Release](https://img.shields.io/github/release/atc0005/dnsc.svg?style=flat-square)](https://github.com/atc0005/dnsc/releases/latest)
[![GoDoc](https://godoc.org/github.com/atc0005/dnsc?status.svg)](https://godoc.org/github.com/atc0005/dnsc)
![Validate Codebase](https://github.com/atc0005/dnsc/workflows/Validate%20Codebase/badge.svg)
![Validate Docs](https://github.com/atc0005/dnsc/workflows/Validate%20Docs/badge.svg)

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
    - [Use config file for DNS servers list and query types](#use-config-file-for-dns-servers-list-and-query-types)
      - [Lots of output](#lots-of-output)
      - [Short output](#short-output)
    - [Specify DNS servers list via flags](#specify-dns-servers-list-via-flags)
    - [Ignore errors, query all servers](#ignore-errors-query-all-servers)
  - [Inspiration](#inspiration)
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

    # https://blog.cloudflare.com/announcing-1111/
    "1.1.1.1",
]

dns_query_types = [
    "A",
    "AAAA",
    "MX",
    "CNAME",
]
```

See the [`config.example.toml`](config.example.toml) file for a starting point.

### Use config file for DNS servers list and query types

#### Lots of output

```ShellSession
$ dnsc -config-file ./config.example.toml -q www.yahoo.com

  INFO[0000] User-specified config file provided, will attempt to load it
  INFO[0000] Trying to load config file "./config.example.toml"
  INFO[0000] Config file successfully loaded config_file=./config.example.toml


Server            Query            Type     Answers                                                                                                                                                       TTL
---               ---              ---      ---                                                                                                                                                           ---
1.1.1.1           www.yahoo.com    CNAME    atsv2-fp-shed.wg1.b.yahoo.com. (CNAME)                                                                                                                        1082
1.1.1.1           www.yahoo.com    A        atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 72.30.35.9 (A), 72.30.35.10 (A), 98.138.219.231 (A), 98.138.219.232 (A)                                               1082, 41, 41, 41, 41
1.1.1.1           www.yahoo.com    AAAA     atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 2001:4998:58:1836::11 (AAAA), 2001:4998:44:41d::3 (AAAA), 2001:4998:44:41d::4 (AAAA), 2001:4998:58:1836::10 (AAAA)    1082, 33, 33, 33, 33
1.1.1.1           www.yahoo.com    MX       atsv2-fp-shed.wg1.b.yahoo.com. (CNAME)                                                                                                                        1082
208.67.220.220    www.yahoo.com    A        atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 72.30.35.9 (A), 72.30.35.10 (A), 98.138.219.231 (A), 98.138.219.232 (A)                                               1137, 38, 38, 38, 38
208.67.220.220    www.yahoo.com    MX       atsv2-fp-shed.wg1.b.yahoo.com. (CNAME)                                                                                                                        1047
208.67.220.220    www.yahoo.com    AAAA     atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 2001:4998:58:1836::10 (AAAA), 2001:4998:58:1836::11 (AAAA), 2001:4998:44:41d::3 (AAAA), 2001:4998:44:41d::4 (AAAA)    1047, 60, 60, 60, 60
208.67.220.220    www.yahoo.com    CNAME    atsv2-fp-shed.wg1.b.yahoo.com. (CNAME)                                                                                                                        679
208.67.222.222    www.yahoo.com    MX       atsv2-fp-shed.wg1.b.yahoo.com. (CNAME)                                                                                                                        679
208.67.222.222    www.yahoo.com    AAAA     atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 2001:4998:58:1836::10 (AAAA), 2001:4998:58:1836::11 (AAAA), 2001:4998:44:41d::3 (AAAA), 2001:4998:44:41d::4 (AAAA)    1047, 60, 60, 60, 60
208.67.222.222    www.yahoo.com    CNAME    atsv2-fp-shed.wg1.b.yahoo.com. (CNAME)                                                                                                                        1524
208.67.222.222    www.yahoo.com    A        atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 72.30.35.9 (A), 72.30.35.10 (A), 98.138.219.231 (A), 98.138.219.232 (A)                                               679, 55, 55, 55, 55
8.8.4.4           www.yahoo.com    AAAA     atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 2001:4998:58:1836::10 (AAAA), 2001:4998:44:41d::4 (AAAA), 2001:4998:44:41d::3 (AAAA), 2001:4998:58:1836::11 (AAAA)    569, 9, 9, 9, 9
8.8.4.4           www.yahoo.com    MX       atsv2-fp-shed.wg1.b.yahoo.com. (CNAME)                                                                                                                        1121
8.8.4.4           www.yahoo.com    A        atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 72.30.35.9 (A), 72.30.35.10 (A), 98.138.219.231 (A), 98.138.219.232 (A)                                               1008, 35, 35, 35, 35
8.8.4.4           www.yahoo.com    CNAME    atsv2-fp-shed.wg1.b.yahoo.com. (CNAME)                                                                                                                        806
8.8.8.8           www.yahoo.com    MX       atsv2-fp-shed.wg1.b.yahoo.com. (CNAME)                                                                                                                        927
8.8.8.8           www.yahoo.com    AAAA     atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 2001:4998:58:1836::11 (AAAA), 2001:4998:44:41d::3 (AAAA), 2001:4998:44:41d::4 (AAAA), 2001:4998:58:1836::10 (AAAA)    924, 23, 23, 23, 23
8.8.8.8           www.yahoo.com    CNAME    atsv2-fp-shed.wg1.b.yahoo.com. (CNAME)                                                                                                                        903
8.8.8.8           www.yahoo.com    A        atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 72.30.35.9 (A), 72.30.35.10 (A), 98.138.219.231 (A), 98.138.219.232 (A)                                               576, 38, 38, 38, 38
```

#### Short output

```ShellSession
$ dnsc -config-file ./config.example.toml -q www.penzoil.com

  INFO[0000] User-specified config file provided, will attempt to load it
  INFO[0000] Trying to load config file "./config.example.toml"
  INFO[0000] Config file successfully loaded config_file=./config.example.toml


Server     Query              Type    Answers                       TTL
---        ---                ---     ---                           ---
8.8.4.4    www.penzoil.com    A       65.52.64.201 (A)              751
1.1.1.1    www.penzoil.com    A       65.52.64.201 (A)              753
8.8.8.8    www.penzoil.com    A       65.52.64.201 (A)              899
1.1.1.1    www.penzoil.com    AAAA    no records found for query
```

Here the process bailed when the first `AAAA` request (due to config file
setting) failed.

### Specify DNS servers list via flags

You can also specify the DNS servers via CLI flags, though it is a bit more verbose:

```ShellSession
$ dnsc -ds 8.8.8.8 -ds 8.8.4.4 -ds 208.67.220.220 -ds 208.67.222.222 -q www.yahoo.com

  INFO[0000] User-specified config file not provided
  INFO[0000] Trying to load config file "T:\\github\\dnsc\\config.toml"
  WARN[0000] Config file "T:\\github\\dnsc\\config.toml" not found or unable to load
  INFO[0000] Trying to load config file "C:\\Users\\adam\\AppData\\Roaming\\dnsc\\config.toml"
  INFO[0000] Config file successfully loaded config_file=C:\Users\adam\AppData\Roaming\dnsc\config.toml


Server            Query            Type     Answers                                                                                                                                                       TTL
---               ---              ---      ---                                                                                                                                                           ---
208.67.220.220    www.yahoo.com    AAAA     atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 2001:4998:58:1836::10 (AAAA), 2001:4998:58:1836::11 (AAAA), 2001:4998:44:41d::3 (AAAA), 2001:4998:44:41d::4 (AAAA)    1339, 11, 11, 11, 11
208.67.220.220    www.yahoo.com    A        atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 72.30.35.9 (A), 72.30.35.10 (A), 98.138.219.231 (A), 98.138.219.232 (A)                                               1339, 40, 40, 40, 40
208.67.220.220    www.yahoo.com    CNAME    atsv2-fp-shed.wg1.b.yahoo.com. (CNAME)                                                                                                                        1775
208.67.220.220    www.yahoo.com    MX       atsv2-fp-shed.wg1.b.yahoo.com. (CNAME)                                                                                                                        1339
208.67.222.222    www.yahoo.com    CNAME    atsv2-fp-shed.wg1.b.yahoo.com. (CNAME)                                                                                                                        395
208.67.222.222    www.yahoo.com    AAAA     atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 2001:4998:58:1836::10 (AAAA), 2001:4998:58:1836::11 (AAAA), 2001:4998:44:41d::3 (AAAA), 2001:4998:44:41d::4 (AAAA)    1339, 11, 11, 11, 11
208.67.222.222    www.yahoo.com    MX       atsv2-fp-shed.wg1.b.yahoo.com. (CNAME)                                                                                                                        858
208.67.222.222    www.yahoo.com    A        atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 72.30.35.9 (A), 72.30.35.10 (A), 98.138.219.231 (A), 98.138.219.232 (A)                                               858, 35, 35, 35, 35
8.8.4.4           www.yahoo.com    AAAA     atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 2001:4998:58:1836::11 (AAAA), 2001:4998:58:1836::10 (AAAA)                                                            1500, 55, 55
8.8.4.4           www.yahoo.com    MX       atsv2-fp-shed.wg1.b.yahoo.com. (CNAME)                                                                                                                        1297
8.8.4.4           www.yahoo.com    A        atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 72.30.35.9 (A), 72.30.35.10 (A)                                                                                       1581, 18, 18
8.8.4.4           www.yahoo.com    CNAME    atsv2-fp-shed.wg1.b.yahoo.com. (CNAME)                                                                                                                        1580
8.8.8.8           www.yahoo.com    AAAA     atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 2001:4998:58:1836::11 (AAAA), 2001:4998:58:1836::10 (AAAA)                                                            1587, 37, 37
8.8.8.8           www.yahoo.com    MX       atsv2-fp-shed.wg1.b.yahoo.com. (CNAME)                                                                                                                        1483
8.8.8.8           www.yahoo.com    A        atsv2-fp-shed.wg1.b.yahoo.com. (CNAME), 72.30.35.9 (A), 72.30.35.10 (A)                                                                                       1616, 53, 53
8.8.8.8           www.yahoo.com    CNAME    atsv2-fp-shed.wg1.b.yahoo.com. (CNAME)                                                                                                                        1580
```

No errors here since www.yahoo.com had records for all of the requested query
types.

It's also worth pointing out that I have a local copy of the `config.toml`
file which was automatically detected and loaded.

### Ignore errors, query all servers

By default, one failure causes the application to immediately fail and query
results against other servers to be discarded. You can override this behavior
to allow the results to be processed from other servers. This can be useful if
you have one DNS server from the group unreachable or providing invalid
results.

Here we use the short-hand `-ide` flag:

```ShellSession
$ dnsc  -ds 8.8.8.8 -ds 8.8.4.4 -ds 208.67.220.220 -ds 208.67.222.222 -q tacos -ide
  INFO[0000] User-specified config file not provided
  INFO[0000] Trying to load config file "T:\\github\\dnsc\\config.toml"
  WARN[0000] Config file "T:\\github\\dnsc\\config.toml" not found or unable to load
  INFO[0000] Trying to load config file "C:\\Users\\adam\\AppData\\Roaming\\dnsc\\config.toml"
  INFO[0000] Config file successfully loaded config_file=C:\Users\adam\AppData\Roaming\dnsc\config.toml


Server            Query    Type     Answers                       TTL
---               ---      ---      ---                           ---
208.67.220.220    tacos    CNAME    no records found for query
208.67.220.220    tacos    A        no records found for query
208.67.220.220    tacos    MX       no records found for query
208.67.220.220    tacos    AAAA     no records found for query
208.67.222.222    tacos    CNAME    no records found for query
208.67.222.222    tacos    AAAA     no records found for query
208.67.222.222    tacos    MX       no records found for query
208.67.222.222    tacos    A        no records found for query
8.8.4.4           tacos    AAAA     no records found for query
8.8.4.4           tacos    CNAME    no records found for query
8.8.4.4           tacos    MX       no records found for query
8.8.4.4           tacos    A        no records found for query
8.8.8.8           tacos    MX       no records found for query
8.8.8.8           tacos    A        no records found for query
8.8.8.8           tacos    CNAME    no records found for query
8.8.8.8           tacos    AAAA     no records found for query
```

Here is the www.penzoil.com query example from earlier, but now with the
`-ide` flag set:

```ShellSession
$ dnsc -config-file ./config.example.toml -q www.penzoil.com -ide
  INFO[0000] User-specified config file provided, will attempt to load it
  INFO[0000] Trying to load config file "./config.example.toml"
  INFO[0000] Config file successfully loaded config_file=./config.example.toml


Server            Query              Type     Answers                       TTL
---               ---                ---      ---                           ---
1.1.1.1           www.penzoil.com    MX       no records found for query
1.1.1.1           www.penzoil.com    A        65.52.64.201 (A)              260
1.1.1.1           www.penzoil.com    CNAME    no records found for query
1.1.1.1           www.penzoil.com    AAAA     no records found for query
208.67.220.220    www.penzoil.com    CNAME    no records found for query
208.67.220.220    www.penzoil.com    MX       no records found for query
208.67.220.220    www.penzoil.com    AAAA     no records found for query
208.67.220.220    www.penzoil.com    A        65.52.64.201 (A)              900
208.67.222.222    www.penzoil.com    CNAME    no records found for query
208.67.222.222    www.penzoil.com    A        65.52.64.201 (A)              497
208.67.222.222    www.penzoil.com    AAAA     no records found for query
208.67.222.222    www.penzoil.com    MX       no records found for query
8.8.4.4           www.penzoil.com    A        65.52.64.201 (A)              258
8.8.4.4           www.penzoil.com    MX       no records found for query
8.8.4.4           www.penzoil.com    AAAA     no records found for query
8.8.4.4           www.penzoil.com    CNAME    no records found for query
8.8.8.8           www.penzoil.com    CNAME    no records found for query
8.8.8.8           www.penzoil.com    MX       no records found for query
8.8.8.8           www.penzoil.com    AAAA     no records found for query
8.8.8.8           www.penzoil.com    A        65.52.64.201 (A)              406
```

## Inspiration

This project was inspired by a small script that I used for many years to
quickly spot outliers with DNS servers after moving virtual servers between
subnets. The shell script is usable, but crude. It is provided as-is and can
be found here:

[dns_query.sh](contrib/dns_query.sh)

This project also owes its existence to the `miekg/dns` package and the Black
Hat Go authors whose example code got me started. See the
[References](#references) section for links to both resources. Many thanks to the
authors of both!

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
  - <https://blog.cloudflare.com/announcing-1111/>
