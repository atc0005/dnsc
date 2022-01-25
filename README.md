<!-- omit in toc -->
# dnsc

Submit query against a list of DNS servers and display summary of results

[![Latest Release](https://img.shields.io/github/release/atc0005/dnsc.svg?style=flat-square)](https://github.com/atc0005/dnsc/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/github.com/atc0005/dnsc.svg)](https://pkg.go.dev/github.com/atc0005/dnsc)
[![Validate Codebase](https://github.com/atc0005/dnsc/workflows/Validate%20Codebase/badge.svg)](https://github.com/atc0005/dnsc/actions?query=workflow%3A%22Validate+Codebase%22)
[![Validate Docs](https://github.com/atc0005/dnsc/workflows/Validate%20Docs/badge.svg)](https://github.com/atc0005/dnsc/actions?query=workflow%3A%22Validate+Docs%22)
[![Lint and Build using Makefile](https://github.com/atc0005/dnsc/workflows/Lint%20and%20Build%20using%20Makefile/badge.svg)](https://github.com/atc0005/dnsc/actions?query=workflow%3A%22Lint+and+Build+using+Makefile%22)
[![Quick Validation](https://github.com/atc0005/dnsc/workflows/Quick%20Validation/badge.svg)](https://github.com/atc0005/dnsc/actions?query=workflow%3A%22Quick+Validation%22)

<!-- omit in toc -->
## Table of Contents

- [Project home](#project-home)
- [Overview](#overview)
- [Features](#features)
  - [Current](#current)
  - [Planned](#planned)
- [Changelog](#changelog)
- [Requirements](#requirements)
  - [Building source code](#building-source-code)
  - [Running](#running)
- [Installation](#installation)
  - [From source](#from-source)
  - [Using release binaries](#using-release-binaries)
- [Configuration](#configuration)
  - [Precedence](#precedence)
  - [Query types supported](#query-types-supported)
  - [Service Location (SRV) Protocol "shortcuts"](#service-location-srv-protocol-shortcuts)
  - [Command-line arguments](#command-line-arguments)
  - [Configuration file](#configuration-file)
- [Examples](#examples)
  - [Our config file](#our-config-file)
  - [Flags only, no config file](#flags-only-no-config-file)
  - [Use config file for DNS servers list and query types](#use-config-file-for-dns-servers-list-and-query-types)
  - [Specify DNS servers list via flags](#specify-dns-servers-list-via-flags)
  - [Query pointer record (PTR) using IP Address](#query-pointer-record-ptr-using-ip-address)
  - [Query server record (SRV)](#query-server-record-srv)
  - [Query server record (SRV) using SRV protocol keyword (aka, "shortcut")](#query-server-record-srv-using-srv-protocol-keyword-aka-shortcut)
  - [Force exit on first DNS error](#force-exit-on-first-dns-error)
  - [Use single-line summary output format](#use-single-line-summary-output-format)
- [Inspiration](#inspiration)
- [References](#references)

## Project home

See [our GitHub repo][repo-url] for the latest code, to file an issue or
submit improvements for review and potential inclusion into the project.

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
- Multiple [query types supported](#query-types-supported)

- Multiple [Service Location (SRV) Protocol "shortcuts" supported](#service-location-srv-protocol-shortcuts)

- User configurable logging levels

- User configurable logging format

- User configurable results summary output format

- User configurable query timeout

### Planned

See [our GitHub repo][repo-url] for planned future work.

## Changelog

See the [`CHANGELOG.md`](CHANGELOG.md) file for the changes associated with
each release of this application. Changes that have been merged to `master`,
but not yet an official release may also be noted in the file under the
`Unreleased` section. A helpful link to the Git commit history since the last
official release is also provided for further review.

## Requirements

The following is a loose guideline. Other combinations of Go and operating
systems for building and running tools from this repo may work, but have not
been tested.

### Building source code

- Go
  - see this project's `go.mod` file for *preferred* version
  - this project tests against [officially supported Go
    releases][go-supported-releases]
    - the most recent stable release (aka, "stable")
    - the prior, but still supported release (aka, "oldstable")
- GCC
  - if building with custom options (as the provided `Makefile` does)
- `make`
  - if using the provided `Makefile`

### Running

- Windows 10
- Ubuntu Linux 18.04+
- Red Hat Enterprise Linux 7+

## Installation

### From source

1. [Download][go-docs-download] Go
1. [Install][go-docs-install] Go
1. Clone the repo
   1. `cd /tmp`
   1. `git clone https://github.com/atc0005/dnsc`
   1. `cd dnsc`
1. Install dependencies (optional)
   - for Ubuntu Linux
     - `sudo apt-get install make gcc`
   - for CentOS Linux
     1. `sudo yum install make gcc`
1. Build
   - for current operating system
     - `go build -mod=vendor ./cmd/dnsc/`
       - *forces build to use bundled dependencies in top-level `vendor`
         folder*
   - for all supported platforms (where `make` is installed)
      - `make all`
   - for Windows
      - `make windows`
   - for Linux
     - `make linux`
1. Copy the applicable binary to whatever systems needs to run it
   - if using `Makefile`: look in `/tmp/dnsc/release_assets/dnsc/`
   - if using `go build`: look in `/tmp/dnsc/`

### Using release binaries

1. Download the [latest
   release](https://github.com/atc0005/dnsc/releases/latest) binaries
1. Deploy
   - Place `dnsc` in a location of your choice
     - e.g., `/usr/local/bin/`

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

### Query types supported

These are the record types currently supported:

- `CNAME`
- `A`
- `AAAA`
- `MX`
- `PTR`
- `SRV`

Other types will be added as I encounter a need for them, or as requested.

### Service Location (SRV) Protocol "shortcuts"

These are the keywords currently supported along with an example of the query
string used for a user-provided query string of `example.com`. Others will be
added as I encounter a need for them, or as requested.

These "shortcuts" are entirely optional. For example, you may still specify
the `srv` query type and provide the `_ldap._tcp.dc._msdcs.example.com` query
string manually to get the same effect.

| User-specified query string | Keyword      | Query string submitted to servers  |
| --------------------------- | ------------ | ---------------------------------- |
| `example.com`               | `msdcs`      | `_ldap._tcp.dc._msdcs.example.com` |
| `example.com`               | `kerberos`   | `_kerberos._tcp.example.com`       |
| `example.com`               | `xmppsrv`    | `_xmpp-server._tcp.example.com`    |
| `example.com`               | `xmppclient` | `_xmpp-client._tcp.example.com`    |
| `example.com`               | `sip`        | `_sip._tcp.example.com`            |

### Command-line arguments

- Flags marked as **`required`** must be set via CLI flag *or* within a
  TOML-formatted configuration file.
- Flags *not* marked as required are for settings where a useful default is
  already defined.

| Flag                      | Required | Default        | Repeat  | Possible                                                       | Description                                                                                                                                                                                                                                                                                                                                                    |
| ------------------------- | -------- | -------------- | ------- | -------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `h`, `help`               | No       | `false`        | No      | `h`, `help`                                                    | Show Help text along with the list of supported flags.                                                                                                                                                                                                                                                                                                         |
| `ds`, `dns-server`        | **Yes**  | *empty string* | **Yes** | *one valid IP Address per flag invocation*                     | DNS server to submit query against. This flag may be repeated for each additional DNS server to query.                                                                                                                                                                                                                                                         |
| `cf`, `config-file`       | **Yes**  | *empty string* | No      | *valid file name characters*                                   | Full path to TOML-formatted configuration file. See [`config.example.toml`](config.example.toml) for a starter template.                                                                                                                                                                                                                                       |
| `v`, `version`            | No       | `false`        | No      | `v`, `version`                                                 | Whether to display application version and then immediately exit application.                                                                                                                                                                                                                                                                                  |
| `def`, `dns-errors-fatal` | No       | `false`        | No      | `def`, `dns-errors-fatal`                                      | Whether DNS-related errors should force this application to immediately exit.                                                                                                                                                                                                                                                                                  |
| `ot`, `omit-timestamp`    | No       | `false`        | No      | `ot`, `omit-timestamp`                                         | Whether the date & time for when the output is generated is omitted from the results output.                                                                                                                                                                                                                                                                   |
| `q`, `query`              | **Yes**  | *empty string* | No      | *any valid FQDN string*                                        | Fully-qualified system to lookup from all provided DNS servers.                                                                                                                                                                                                                                                                                                |
| `sp`, `srv-protocol`      | No       | *empty list*   | **Yes** | [supported keywords](#service-location-srv-protocol-shortcuts) | Service Location (SRV) protocols associated with a given domain name as the query string. For example, `msdcs` can be specified as the SRV record protocol along with `example.com` as the query string to search DNS for `_ldap._tcp.dc._msdcs.example.com`. This flag may be repeated for each additional SRV protocol that you wish to request records for. |
| `ll`, `log-level`         | No       | `info`         | No      | `fatal`, `error`, `warn`, `info`, `debug`                      | Log message priority filter. Log messages with a lower level are ignored.                                                                                                                                                                                                                                                                                      |
| `lf`, `log-format`        | No       | `text`         | No      | `cli`, `json`, `logfmt`, `text`, `discard`                     | Use the specified `apex/log` package "handler" to output log messages in that handler's format.                                                                                                                                                                                                                                                                |
| `ro`, `results-output`    | No       | `multi-line`   | No      | `multi-line`, `single-line`                                    | Specifies whether the results summary output is composed of a single comma-separated line of records for a query, or whether the records are returned one per line.                                                                                                                                                                                            |
| `t`, `type`               | No       | `A`            | **Yes** | [supported types](#query-types-supported)                      | DNS query type to use when submitting a DNS query to each provided server. This flag may be repeated for each additional DNS record type you wish to request.                                                                                                                                                                                                  |
| `to`, `timeout`           | No       | `10`           | No      | *any positive whole number*                                    | Maximum number of seconds allowed for a DNS query to take before timing out.                                                                                                                                                                                                                                                                                   |

### Configuration file

Configuration file settings have the lowest priority and are overridden by
settings specified in other configuration sources, except for default values.
See the [Command-line Arguments](#command-line-arguments) table for more
information, including the available values for the listed configuration
settings.

| Flag Name          | Config file Setting Name | Notes                                                                              |
| ------------------ | ------------------------ | ---------------------------------------------------------------------------------- |
| `dns-server`       | `dns_servers`            | [Multi-line array](https://github.com/toml-lang/toml#user-content-array)           |
| `query`            | `query`                  | While supported, having a fixed query in the config file is not a normal use case. |
| `dns-errors-fatal` | `dns_errors_fatal`       | Opt-in setting. Useful to leave as-is for most use cases.                          |
| `log-level`        | `log_level`              |                                                                                    |
| `log-format`       | `log_format`             |                                                                                    |
| `results-output`   | `results_output`         |                                                                                    |
| `omit-timestamp`   | `omit_timestamp`         |                                                                                    |
| `type`             | `dns_request_types`      | [Multi-line array](https://github.com/toml-lang/toml#user-content-array)           |
| `srv-protocol`     | `dns_srv_protocols`      | [Multi-line array](https://github.com/toml-lang/toml#user-content-array)           |
| `timeout`          | `timeout`                |                                                                                    |

The [`config.example.toml`](config.example.toml) file is intended as a
starting point for your own `config.toml` configuration file and attempts to
illustrate working values for the available command-line flags.

Once reviewed and potentially adjusted, your copy of the `config.toml` file
can be placed in one of the following locations to be automatically detected
and used by this application:

- alongside the `dnsc` (or `dnsc.exe`) binary (as `config.toml`)
- at `$HOME/.config/dnsc/config.toml` on a UNIX-like system (e.g., Linux
  distro, Mac)
- at `C:\Users\YOUR_USERNAME\AppData\dnsc\config.toml` on a Windows system

Feel free to place the file wherever you like and refer to it using the `-cf`
(short flag name) or `-config-file` (full-length flag name). See the
[Examples](#examples) and [Command-line arguments](#command-line-arguments)
sections for usage details.

## Examples

### Our config file

These examples assume that the referenced `config.example.toml` file has the
following contents:

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

results_output = "multi-line"
```

See the [Configuration file](#configuration-file) section for additional
information, including supported locations for this file.

### Flags only, no config file

Here the `dnsc` application attempted to load a config file using the
locations mentioned in the [Configuration file](#configuration-file) section,
but one was not found:

```ShellSession
$ dnsc -ds 8.8.8.8 -ds 8.8.4.4 -ds 208.67.220.220 -ds 208.67.222.222 -q www.yahoo.com
  INFO[0000] User-specified config file not provided
  INFO[0000] Trying to load config file "/mnt/t/github/dnsc/config.toml"
  WARN[0000] Config file "/mnt/t/github/dnsc/config.toml" not found or unable to load
  INFO[0000] Trying to load config file "/home/ubuntu/.config/dnsc/config.toml"
  WARN[0000] Config file "/home/ubuntu/.config/dnsc/config.toml" not found or unable to load
  WARN[0000] Failed to load config files, relying only on provided flag settings


Server            RTT     Query            Query Type    Answer                          Answer Type    TTL
---               ---     ---              ---           ---                             ---            ---
208.67.220.220    20ms    www.yahoo.com    A             new-fp-shed.wg1.b.yahoo.com.    CNAME          21
208.67.220.220    20ms    www.yahoo.com    A             74.6.143.25                     A              33
208.67.220.220    20ms    www.yahoo.com    A             74.6.143.26                     A              33
208.67.220.220    20ms    www.yahoo.com    A             74.6.231.20                     A              33
208.67.220.220    20ms    www.yahoo.com    A             74.6.231.21                     A              33
208.67.222.222    20ms    www.yahoo.com    A             new-fp-shed.wg1.b.yahoo.com.    CNAME          54
208.67.222.222    20ms    www.yahoo.com    A             74.6.143.25                     A              29
208.67.222.222    20ms    www.yahoo.com    A             74.6.143.26                     A              29
208.67.222.222    20ms    www.yahoo.com    A             74.6.231.20                     A              29
208.67.222.222    20ms    www.yahoo.com    A             74.6.231.21                     A              29
8.8.4.4           7ms     www.yahoo.com    A             new-fp-shed.wg1.b.yahoo.com.    CNAME          36
8.8.4.4           7ms     www.yahoo.com    A             74.6.143.25                     A              45
8.8.4.4           7ms     www.yahoo.com    A             74.6.143.26                     A              45
8.8.4.4           7ms     www.yahoo.com    A             74.6.231.20                     A              45
8.8.4.4           7ms     www.yahoo.com    A             74.6.231.21                     A              45
8.8.8.8           8ms     www.yahoo.com    A             new-fp-shed.wg1.b.yahoo.com.    CNAME          25
8.8.8.8           8ms     www.yahoo.com    A             74.6.143.25                     A              22
8.8.8.8           8ms     www.yahoo.com    A             74.6.143.26                     A              22
8.8.8.8           8ms     www.yahoo.com    A             74.6.231.20                     A              22
8.8.8.8           8ms     www.yahoo.com    A             74.6.231.21                     A              22
```

Since we did not specify a record query type, only `A` records were requested
in the query to each DNS server. Since we have recursion enabled and the
`dnsc` application does not exclude them, the initial `CNAME` record found by
each query is also listed in the `Answers` section above for each DNS server
that had it.

### Use config file for DNS servers list and query types

As can be seen below, this example produces quite a bit of output.

```ShellSession
$ dnsc --config-file ./config.example.toml -q www.yahoo.com
  INFO[0000] User-specified config file provided, will attempt to load it
  INFO[0000] Trying to load config file "./config.example.toml"
  INFO[0000] Config file successfully loaded config_file=./config.example.toml


Server            RTT     Query            Type     Answers                         TTL
---               ---     ---              ---      ---                             ---
1.1.1.1           7ms     www.yahoo.com    CNAME    new-fp-shed.wg1.b.yahoo.com.    8
1.1.1.1           7ms     www.yahoo.com    A        74.6.143.25                     8
1.1.1.1           7ms     www.yahoo.com    A        74.6.143.26                     8
1.1.1.1           7ms     www.yahoo.com    A        74.6.231.20                     8
1.1.1.1           7ms     www.yahoo.com    A        74.6.231.21                     8
1.1.1.1           51ms    www.yahoo.com    CNAME    new-fp-shed.wg1.b.yahoo.com.    60
1.1.1.1           7ms     www.yahoo.com    CNAME    new-fp-shed.wg1.b.yahoo.com.    12
1.1.1.1           7ms     www.yahoo.com    AAAA     2001:4998:124:1507::f000        12
1.1.1.1           7ms     www.yahoo.com    AAAA     2001:4998:124:1507::f001        12
1.1.1.1           7ms     www.yahoo.com    AAAA     2001:4998:44:3507::8000         12
1.1.1.1           7ms     www.yahoo.com    AAAA     2001:4998:44:3507::8001         12
1.1.1.1           73ms    www.yahoo.com    CNAME    new-fp-shed.wg1.b.yahoo.com.    60
208.67.220.220    21ms    www.yahoo.com    CNAME    new-fp-shed.wg1.b.yahoo.com.    19
208.67.220.220    21ms    www.yahoo.com    AAAA     2001:4998:44:3507::8001         16
208.67.220.220    21ms    www.yahoo.com    AAAA     2001:4998:124:1507::f000        16
208.67.220.220    21ms    www.yahoo.com    AAAA     2001:4998:124:1507::f001        16
208.67.220.220    21ms    www.yahoo.com    AAAA     2001:4998:44:3507::8000         16
208.67.220.220    23ms    www.yahoo.com    CNAME    new-fp-shed.wg1.b.yahoo.com.    42
208.67.220.220    22ms    www.yahoo.com    CNAME    new-fp-shed.wg1.b.yahoo.com.    20
208.67.220.220    22ms    www.yahoo.com    CNAME    new-fp-shed.wg1.b.yahoo.com.    19
208.67.220.220    22ms    www.yahoo.com    A        74.6.143.25                     58
208.67.220.220    22ms    www.yahoo.com    A        74.6.143.26                     58
208.67.220.220    22ms    www.yahoo.com    A        74.6.231.20                     58
208.67.220.220    22ms    www.yahoo.com    A        74.6.231.21                     58
208.67.222.222    22ms    www.yahoo.com    CNAME    new-fp-shed.wg1.b.yahoo.com.    25
208.67.222.222    22ms    www.yahoo.com    A        74.6.143.25                     50
208.67.222.222    22ms    www.yahoo.com    A        74.6.143.26                     50
208.67.222.222    22ms    www.yahoo.com    A        74.6.231.20                     50
208.67.222.222    22ms    www.yahoo.com    A        74.6.231.21                     50
208.67.222.222    21ms    www.yahoo.com    CNAME    new-fp-shed.wg1.b.yahoo.com.    19
208.67.222.222    21ms    www.yahoo.com    AAAA     2001:4998:44:3507::8001         16
208.67.222.222    21ms    www.yahoo.com    AAAA     2001:4998:124:1507::f000        16
208.67.222.222    21ms    www.yahoo.com    AAAA     2001:4998:124:1507::f001        16
208.67.222.222    21ms    www.yahoo.com    AAAA     2001:4998:44:3507::8000         16
208.67.222.222    21ms    www.yahoo.com    CNAME    new-fp-shed.wg1.b.yahoo.com.    20
208.67.222.222    23ms    www.yahoo.com    CNAME    new-fp-shed.wg1.b.yahoo.com.    19
8.8.4.4           7ms     www.yahoo.com    CNAME    new-fp-shed.wg1.b.yahoo.com.    34
8.8.4.4           8ms     www.yahoo.com    CNAME    new-fp-shed.wg1.b.yahoo.com.    52
8.8.4.4           8ms     www.yahoo.com    CNAME    new-fp-shed.wg1.b.yahoo.com.    34
8.8.4.4           8ms     www.yahoo.com    AAAA     2001:4998:44:3507::8001         1
8.8.4.4           8ms     www.yahoo.com    AAAA     2001:4998:124:1507::f000        1
8.8.4.4           8ms     www.yahoo.com    AAAA     2001:4998:44:3507::8000         1
8.8.4.4           8ms     www.yahoo.com    AAAA     2001:4998:124:1507::f001        1
8.8.4.4           7ms     www.yahoo.com    CNAME    new-fp-shed.wg1.b.yahoo.com.    33
8.8.4.4           7ms     www.yahoo.com    A        74.6.143.25                     3
8.8.4.4           7ms     www.yahoo.com    A        74.6.143.26                     3
8.8.4.4           7ms     www.yahoo.com    A        74.6.231.20                     3
8.8.4.4           7ms     www.yahoo.com    A        74.6.231.21                     3
8.8.8.8           7ms     www.yahoo.com    CNAME    new-fp-shed.wg1.b.yahoo.com.    34
8.8.8.8           7ms     www.yahoo.com    CNAME    new-fp-shed.wg1.b.yahoo.com.    27
8.8.8.8           7ms     www.yahoo.com    AAAA     2001:4998:124:1507::f000        9
8.8.8.8           7ms     www.yahoo.com    AAAA     2001:4998:124:1507::f001        9
8.8.8.8           7ms     www.yahoo.com    AAAA     2001:4998:44:3507::8000         9
8.8.8.8           7ms     www.yahoo.com    AAAA     2001:4998:44:3507::8001         9
8.8.8.8           7ms     www.yahoo.com    CNAME    new-fp-shed.wg1.b.yahoo.com.    52
8.8.8.8           7ms     www.yahoo.com    A        74.6.143.25                     1
8.8.8.8           7ms     www.yahoo.com    A        74.6.143.26                     1
8.8.8.8           7ms     www.yahoo.com    A        74.6.231.20                     1
8.8.8.8           7ms     www.yahoo.com    A        74.6.231.21                     1
8.8.8.8           7ms     www.yahoo.com    CNAME    new-fp-shed.wg1.b.yahoo.com.    52
```

### Specify DNS servers list via flags

You can also specify the DNS servers via CLI flags, though it is a bit more verbose:

```ShellSession
$ dnsc -ds 8.8.8.8 -ds 8.8.4.4 -ds 208.67.220.220 -ds 208.67.222.222 -q www.yahoo.com
  INFO[0000] User-specified config file not provided
  INFO[0000] Trying to load config file "/mnt/t/github/dnsc/config.toml"
  WARN[0000] Config file "/mnt/t/github/dnsc/config.toml" not found or unable to load
  INFO[0000] Trying to load config file "/home/ubuntu/.config/dnsc/config.toml"
  WARN[0000] Config file "/home/ubuntu/.config/dnsc/config.toml" not found or unable to load
  WARN[0000] Failed to load config files, relying only on provided flag settings


Server            RTT     Query            Query Type    Answer                          Answer Type    TTL
---               ---     ---              ---           ---                             ---            ---
208.67.220.220    20ms    www.yahoo.com    A             new-fp-shed.wg1.b.yahoo.com.    CNAME          43
208.67.220.220    20ms    www.yahoo.com    A             74.6.143.25                     A              48
208.67.220.220    20ms    www.yahoo.com    A             74.6.143.26                     A              48
208.67.220.220    20ms    www.yahoo.com    A             74.6.231.20                     A              48
208.67.220.220    20ms    www.yahoo.com    A             74.6.231.21                     A              48
208.67.222.222    20ms    www.yahoo.com    A             new-fp-shed.wg1.b.yahoo.com.    CNAME          54
208.67.222.222    20ms    www.yahoo.com    A             74.6.143.24                     A              42
208.67.222.222    20ms    www.yahoo.com    A             74.6.231.19                     A              42
8.8.4.4           7ms     www.yahoo.com    A             new-fp-shed.wg1.b.yahoo.com.    CNAME          44
8.8.4.4           7ms     www.yahoo.com    A             74.6.143.25                     A              44
8.8.4.4           7ms     www.yahoo.com    A             74.6.143.26                     A              44
8.8.4.4           7ms     www.yahoo.com    A             74.6.231.20                     A              44
8.8.4.4           7ms     www.yahoo.com    A             74.6.231.21                     A              44
8.8.8.8           7ms     www.yahoo.com    A             new-fp-shed.wg1.b.yahoo.com.    CNAME          44
8.8.8.8           7ms     www.yahoo.com    A             74.6.143.25                     A              44
8.8.8.8           7ms     www.yahoo.com    A             74.6.143.26                     A              44
8.8.8.8           7ms     www.yahoo.com    A             74.6.231.20                     A              44
8.8.8.8           7ms     www.yahoo.com    A             74.6.231.21                     A              44
```

No errors here since www.yahoo.com had records for all of the requested query
types.

### Query pointer record (PTR) using IP Address

Not shown is the initial query for `www.yahoo.com` which resolved to multiple
`A` records. I picked one of the IP Addresses from the list and used it in the
example below.

```ShellSession

$ dnsc --ds 8.8.8.8 --q 74.6.143.25 -t ptr
  INFO[0000] User-specified config file not provided
  INFO[0000] Trying to load config file "/mnt/t/github/dnsc/config.toml"
  WARN[0000] Config file "/mnt/t/github/dnsc/config.toml" not found or unable to load
  INFO[0000] Trying to load config file "/home/ubuntu/.config/dnsc/config.toml"
  WARN[0000] Config file "/home/ubuntu/.config/dnsc/config.toml" not found or unable to load
  WARN[0000] Failed to load config files, relying only on provided flag settings


Server     RTT    Query          Query Type    Answer                                             Answer Type    TTL
---        ---    ---            ---           ---                                                ---            ---
8.8.8.8    7ms    74.6.143.25    PTR           media-router-fp73.prod.media.vip.bf1.yahoo.com.    PTR            186
```

### Query server record (SRV)

In this example, we query for available XMPP servers for `conversations.im`.

```ShellSession

$ dnsc --ds 8.8.8.8 --q "_xmpp-client._tcp.conversations.im" -t srv
  INFO[0000] User-specified config file not provided
  INFO[0000] Trying to load config file "/mnt/t/github/dnsc/config.toml"
  WARN[0000] Config file "/mnt/t/github/dnsc/config.toml" not found or unable to load
  INFO[0000] Trying to load config file "/home/ubuntu/.config/dnsc/config.toml"
  WARN[0000] Config file "/home/ubuntu/.config/dnsc/config.toml" not found or unable to load
  WARN[0000] Failed to load config files, relying only on provided flag settings


Server     RTT    Query                                 Query Type    Answer                     Answer Type    TTL
---        ---    ---                                   ---           ---                        ---            ---
8.8.8.8    7ms    _xmpp-client._tcp.conversations.im    SRV           xmpp.conversations.im.     SRV            2533
8.8.8.8    7ms    _xmpp-client._tcp.conversations.im    SRV           xmpps.conversations.im.    SRV            2533
```

### Query server record (SRV) using SRV protocol keyword (aka, "shortcut")

In this example, we query for available XMPP servers for `conversations.im`
using a SRV protocol keyword or "shortcut". This makes the query a little less
verbose to type. You may optionally omit `-t srv` from this example and you
will receive the same result; if specifying a SRV protocol keyword, this
application assumes that you wish to submit a SRV record query.

```ShellSession
$ dnsc --ds 8.8.8.8 --q "conversations.im" -sp "xmppclient" -t srv
  INFO[0000] User-specified config file not provided
  INFO[0000] Trying to load config file "/mnt/t/github/dnsc/config.toml"
  WARN[0000] Config file "/mnt/t/github/dnsc/config.toml" not found or unable to load
  INFO[0000] Trying to load config file "/home/ubuntu/.config/dnsc/config.toml"
  WARN[0000] Config file "/home/ubuntu/.config/dnsc/config.toml" not found or unable to load
  WARN[0000] Failed to load config files, relying only on provided flag settings


Server     RTT    Query                                 Query Type    Answer                     Answer Type    TTL
---        ---    ---                                   ---           ---                        ---            ---
8.8.8.8    8ms    _xmpp-client._tcp.conversations.im    SRV           xmpps.conversations.im.    SRV            2837
8.8.8.8    8ms    _xmpp-client._tcp.conversations.im    SRV           xmpp.conversations.im.     SRV            2837
```

As with other record/query types, you may also mix several together. Here we
specify the SRV protocol as before, but leave out the explicit `-t srv`
(relying on implicit inclusion) and also specify that we wish to query for
available `A` records.

```ShellSession
$ dnsc --ds 8.8.8.8 --q "conversations.im" -sp "xmppclient" -t a
  INFO[0000] User-specified config file not provided
  INFO[0000] Trying to load config file "/mnt/t/github/dnsc/config.toml"
  WARN[0000] Config file "/mnt/t/github/dnsc/config.toml" not found or unable to load
  INFO[0000] Trying to load config file "/home/ubuntu/.config/dnsc/config.toml"
  WARN[0000] Config file "/home/ubuntu/.config/dnsc/config.toml" not found or unable to load
  WARN[0000] Failed to load config files, relying only on provided flag settings


Server     RTT      Query                                 Query Type    Answer                     Answer Type    TTL
---        ---      ---                                   ---           ---                        ---            ---
8.8.8.8    27ms     _xmpp-client._tcp.conversations.im    SRV           xmpp.conversations.im.     SRV            3599
8.8.8.8    27ms     _xmpp-client._tcp.conversations.im    SRV           xmpps.conversations.im.    SRV            3599
8.8.8.8    231ms    conversations.im                      A             78.47.177.120              A              3599
```

### Force exit on first DNS error

As of `v0.3.0`, this application ignores query errors by default in order to
allow the results to be processed from all servers. This can be useful if you
have one DNS server from the group unreachable or providing invalid results.
If however you wish to restore the previous behavior you can specify either
the `dns-errors-fatal` or `def` flags, or set `dns_errors_fatal = true` in the
`config.toml` config file. The example below assumes that you've set
`dns_errors_fatal = true` in the `config.toml` config file.

Example:

```ShellSession
$ dnsc -config-file ./config.example.toml -q www.penzoil.com
  INFO[0000] User-specified config file provided, will attempt to load it
  INFO[0000] Trying to load config file "./config.example.toml"
  INFO[0000] Config file successfully loaded config_file=./config.example.toml


Server     RTT     Query              Type    Answers                       TTL
---        ---     ---                ---     ---                           ---
8.8.4.4    33ms    www.penzoil.com    A       65.52.64.201 (A)              899
8.8.4.4    8ms     www.penzoil.com    AAAA    no records found for query
```

Here the process bailed when the first `AAAA` request failed. This is due to
the `dns_errors_fatal = true` config file setting noted earlier.

### Use single-line summary output format

This is the previous default output format. This format reads well when the
result set is small, but may be harder to read when many records are returned
for a query type (e.g., large pool of front-end web servers).

```ShellSession
$ dnsc --ds 8.8.8.8 --q "conversations.im" -sp "xmppclient" -t a --ro single-line
  INFO[0000] User-specified config file not provided
  INFO[0000] Trying to load config file "/mnt/t/github/dnsc/config.toml"
  WARN[0000] Config file "/mnt/t/github/dnsc/config.toml" not found or unable to load
  INFO[0000] Trying to load config file "/home/ubuntu/.config/dnsc/config.toml"
  WARN[0000] Config file "/home/ubuntu/.config/dnsc/config.toml" not found or unable to load
  WARN[0000] Failed to load config files, relying only on provided flag settings


Server     RTT      Query                                 Type    Answers                                                        TTL
---        ---      ---                                   ---     ---                                                            ---
8.8.8.8    8ms      _xmpp-client._tcp.conversations.im    SRV     xmpps.conversations.im. (SRV), xmpp.conversations.im. (SRV)    1545, 1545
8.8.8.8    944ms    conversations.im                      A       78.47.177.120 (A)                                              3599
```

Here is an example where this format does not read quite as well:

```ShellSession
$ dnsc --ds 8.8.8.8 -q www.yahoo.com -t a -t aaaa --ro single-line
  INFO[0000] User-specified config file not provided
  INFO[0000] Trying to load config file "/mnt/t/github/dnsc/config.toml"
  WARN[0000] Config file "/mnt/t/github/dnsc/config.toml" not found or unable to load
  INFO[0000] Trying to load config file "/home/ubuntu/.config/dnsc/config.toml"
  WARN[0000] Config file "/home/ubuntu/.config/dnsc/config.toml" not found or unable to load
  WARN[0000] Failed to load config files, relying only on provided flag settings


Server     RTT    Query            Type    Answers                                                                                                                                                                   TTL
---        ---    ---              ---     ---                                                                                                                                                                       ---
8.8.8.8    8ms    www.yahoo.com    AAAA    new-fp-shed.wg1.b.yahoo.com. (CNAME), 2001:4998:44:3507::8000 (AAAA), 2001:4998:44:3507::8001 (AAAA), 2001:4998:124:1507::f001 (AAAA), 2001:4998:124:1507::f000 (AAAA)    5, 5, 5, 5, 5
8.8.8.8    7ms    www.yahoo.com    A       new-fp-shed.wg1.b.yahoo.com. (CNAME), 74.6.143.25 (A), 74.6.143.26 (A), 74.6.231.20 (A), 74.6.231.21 (A)                                                                  39, 40, 40, 40, 40
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

- Definitions
  - <https://www.cloudflare.com/learning/cdn/glossary/round-trip-time-rtt/>
