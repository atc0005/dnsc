# Changelog

## Overview

All notable changes to this project will be documented in this file.

The format is based on [Keep a
Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Please [open an issue](https://github.com/atc0005/dnsc/issues) for any
deviations that you spot; I'm still learning!.

## Types of changes

The following types of changes will be recorded in this file:

- `Added` for new features.
- `Changed` for changes in existing functionality.
- `Deprecated` for soon-to-be removed features.
- `Removed` for now removed features.
- `Fixed` for any bug fixes.
- `Security` in case of vulnerabilities.

## [Unreleased]

- placeholder

## [v0.6.3] - 2022-03-02

### Overview

- Dependency updates
- CI / linting improvements
- built using Go 1.17.7
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.6` to `1.17.7`
  - `miekg/dns`
    - `v1.1.45` to `v1.1.46`
  - `actions/checkout`
    - `v2.4.0` to `v3`
  - `actions/setup-node`
    - `v2.5.1` to `v3`

- (GH-221) Expand linting GitHub Actions Workflow to include `oldstable`,
  `unstable` container images
- (GH-222) Switch Docker image source from Docker Hub to GitHub Container
  Registry (GHCR)

## [v0.6.2] - 2022-01-25

### Overview

- Dependency updates
- built using Go 1.17.6
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.12` to `1.17.6`
    - (GH-218) Update go.mod file, canary Dockerfile to reflect current
      dependencies

## [v0.6.1] - 2021-12-28

### Overview

- Dependency updates
- built using Go 1.16.12
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `miekg/dns`
    - `v1.1.43` to `v1.1.45`
  - `actions/setup-node`
    - `v2.5.0` to `v2.5.1`

## [v0.6.0] - 2021-12-17

### Overview

- Bugfixes
- Output tweak
- Dependency updates
- built using Go 1.16.12
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- (GH-161) Add timestamp to query results output with optional flag & config
  setting to omit if desired

### Changed

- Dependencies
  - `Go`
    - `1.16.10` to `1.16.12`
  - `actions/setup-node`
    - `v2.4.1` to `v2.5.0`

### Fixed

- (GH-206) Deferred logging statement missing filename reference
- (GH-207) Add missing filename to deferred logging call
- (GH-210) Minor typo fixes

## [v0.5.8] - 2021-11-09

### Overview

- Dependency updates
- built using Go 1.16.10
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.8` to `1.16.10`
  - `actions/checkout`
    - `v2.3.4` to `v2.4.0`
  - `actions/setup-node`
    - `v2.4.0` to `v2.4.1`

### Fixed

- (GH-202) False positive `G307: Deferring unsafe method "Close" on type
  "*os.File" (gosec)` linting error

## [v0.5.7] - 2021-09-23

### Overview

- Dependency updates
- built using Go 1.16.8
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.7` to `1.16.8`
  - `pelletier/go-toml`
    - `v1.9.3` to `v1.9.4`

## [v0.5.6] - 2021-08-08

### Overview

- Dependency updates
- built using Go 1.16.7
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.6` to `1.16.7`
  - `actions/setup-node`
    - updated from `v2.2.0` to `v2.4.0`

## [v0.5.5] - 2021-07-18

### Overview

- Dependency updates
- built using Go 1.16.6
  - Statically linked
  - Linux (x86, x64)

### Changed

- dependencies
  - `Go`
    - `1.16.5` to `1.16.6`
  - `actions/setup-node`
    - update `node-version` value to always use latest LTS version instead of
      hard-coded version

## [v0.5.4] - 2021-07-01

### Overview

- Bugfixes
- Output tweak
- built using Go 1.16.5
  - Statically linked
  - Linux (x86, x64)

### Added

- (GH-169) Create "canary" Dockerfile to track stable Go releases, serve as a
  reminder to generate fresh binaries

### Changed

- dependencies
  - (GH-172) `actions/setup-node`
    - `v2.1.5` to `v2.2.0`

- output
  - (GH-171) Exclude unused columns when no records for query are found

### Fixed

- linting
  - (GH-176) var-declaration: should omit type string from declaration of var
    version; it will be inferred from the right-hand side (revive)
  - (GH-173) `golangci/golangci-lint`
    - replace deprecated `golint` linter with `revive`

## [v0.5.3] - 2021-06-23

### Overview

- Dependency updates
- built using Go 1.16.5
  - Statically linked
  - Linux (x86, x64)

### Changed

- `flag.ErrHelp` handling updated
- minimum Go version bumped to 1.15
- dependencies
  - `miekg/dns`
    - `v1.1.41` to `v1.1.43`
  - `pelletier/go-toml`
    - `v1.9.0` to `v1.9.3`

## [v0.5.2] - 2021-04-28

### Overview

- Logging tweak
- built using Go 1.16.3
  - Statically linked
  - Linux (x86, x64)

### Changed

- Change logging level for config file loading status messages from `INFO` to
  `DEBUG`

### Fixed

- Add missing deps to v0.5.1 changelog entry

## [v0.5.1] - 2021-04-05

### Overview

- Bug fixes
- built using Go 1.16.3
  - Statically linked
  - Linux (x86, x64)

### Changed

- dependencies
  - `miekg/dns`
    - `v1.1.40` to `v1.1.41`
  - `pelletier/go-toml`
    - `v1.8.1` to `v1.9.0`

### Fixed

- linting
  - fieldalignment: struct with X pointer bytes could be Y (govet)
  - `golangci/golangci-lint`
    - replace deprecated `maligned` linter with `govet: fieldalignment`
    - replace deprecated `scopelint` linter with `exportloopref`

## [v0.5.0] - 2021-03-08

### Overview

- Add support for SRV record protocol "shortcuts"
- Add new (default) multi-line results summary output format
- Misc bugfixes
- built using Go 1.15.8

### Added

- Add support for SRV record protocol "shortcuts"
  - e.g., allows specifying `msdcs` as protocol keyword and `example.com` as
    the query string to query for available domain controllers instead of
    specifying `_ldap._tcp.dc._msdcs.example.com` as the query string.
- Add new (default) multi-line results summary output format
  - attempts to work around display issues with many results per record type

### Changed

- Default results summary output changed from `single-line` to `multi-line`
  - the prior format can be set persistently via config file or one-off via
    CLI flag
- Modify concurrency implementation to better support future work and help
  with implementing SRV protocol "shortcuts"
  - while this should be an improvement overall, this has not been fully
    tested yet

### Fixed

- Repeating query type flag results in duplicate queries
- Use default consts (which are currently empty strings) instead of actual
  empty strings
  - this was a bug waiting to happen

## [v0.4.0] - 2021-03-04

### Overview

- Add support for additional record types
- Dependency updates
- built using Go 1.15.8

### Added

- Support for `PTR` record queries
- Support for `SRV` record queries

### Changed

- dependencies
  - `miekg/dns`
    - `v1.1.38` to `v1.1.40`
  - `actions/setup-node`
    - `v2.1.4` to `v2.1.5`

## [v0.3.5] - 2021-02-21

### Overview

- Dependency updates
- built using Go 1.15.8

### Changed

- Swap out GoDoc badge for pkg.go.dev badge

- dependencies
  - `go.mod` Go version
    - updated from `1.14` to `1.15`
  - built using Go 1.15.8
    - Statically linked
    - Windows (x86, x64)
    - Linux (x86, x64)
  - `miekg/dns`
    - `v1.1.35` to `v1.1.38`
  - `actions/setup-node`
    - `v2.1.2` to `v2.1.4`

## [v0.3.4] - 2020-11-16

### Changed

- Dependencies
  - built using Go 1.15.5
    - **Statically linked**
    - Windows
      - x86
      - x64
    - Linux
      - x86
      - x64
  - `miekg/dns`
    - `v1.1.31` to `v1.1.35`
  - `actions/checkout`
    - `v2.3.3` to `v2.3.4`

## [v0.3.3] - 2020-10-10

### Added

- Statically linked binary release
  - Built using Go 1.15.2
  - Native Go binaries (no cgo)
  - Windows
    - x86
    - x64
  - Linux
    - x86
    - x64

### Changed

- Dependencies
  - `actions/checkout`
    - `v2.3.2` to `v2.3.3`
  - `actions/setup-node`
    - `v2.1.1` to `v2.1.2`

- Build options updated
  - Add `-trimpath` build flag
  - Explicitly disable cgo
  - Apply `osusergo` and `netgo` build tags
    - help ensure static builds that are not dependent on glibc

### Fixed

- Makefile generates checksums with qualified path
- Makefile build options do not generate static binaries
- Missing shorthand suffix in flags help text
- (Some) getter methods do not appear to return intended default values

## [v0.3.2] - 2020-09-14

### Changed

- Dependencies
  - upgrade `go.mod` Go version
    - `1.13` to `1.14`
  - upgrade `apex/log`
    - `v1.7.0` to `v1.9.0`
  - upgrade `miekg/dns`
    - `v1.1.30` to `v1.1.31`
  - upgrade `mattn/go-colorable`
    - `v0.1.6` to `v0.1.7`
  - upgrade `gopkg.in/check.v1`
    - `v1.0.0-20200227125254-8fa46927fb4f` to
      `v1.0.0-20200902074654-038fdea0a05b`
  - upgrade `actions/checkout`
    - `v2.3.1` to `v2.3.2`

### Fixed

- Explicitly pass `queryTimeout` to goroutine
- Documentation
  - Attempt to clarify config file load behavior
  - Typo
- Misc linting issues surfaced by `golangci-lint` `v1.31.0` release

## [v0.3.1] - 2020-08-04

### Added

- Add new README badges for additional CI workflows
  - each badge also links to the associated workflow results

### Changed

- Dependencies
  - upgrade `apex/log`
    - `v1.6.0` to `v1.7.0`
  - upgrade `miekg/dns`
    - `v1.1.30` to `v1.1.31`

- Linting
  - CI
  - Swap out `go-ci-stable` image tag for `go-ci-lint-only`
    - the `go-ci-lint-only` image is substantially smaller and *should* result
      in faster spin-up times
  - Remove repo-provided copy of `golangci-lint` config file at start of
    linting task in order to force use of Docker container-provided config
    file

## [v0.3.0] - 2020-07-30

### Added

- Add `dns-errors-fatal`, `def` flags to halt application on first dns query
  error

- Use Docker containers from `atc0005/go-ci` project for linting, testing,
  building in place of `actions/setup-go` provided environment
  - "old stable"
    - Go 1.13.x series (currently)
  - "stable"
    - Go 1.14.x series (currently)
  - "unstable"
    - Go 1.15rc1 (currently)

### Changed

- Disable `golangci-lint` default exclusions

- Makefile
  - install latest `golangci-lint` binary (not locked to specific version)

- Ignore query errors by default
  - Replace `ignore-dns-errors`, `ide` flags with inverse `dns-errors-fatal`,
    `def` flags, flip default logic to allow query errors to be ignored by
    default, but force failure on first error if desired.

- Dependencies
  - upgraded `apex/log`
    - `v1.4.0` to `v1.6.0`
  - upgraded `actions/setup-go`
    - `v2.1.0` to `v2.1.1`
  - upgraded `actions/setup-node`
    - `v2.1.0` to `v2.1.1`

### Removed

- Replace `ignore-dns-errors`, `ide` flags with inverse `dns-errors-fatal`,
  `def` flags, flip default logic.

### Fixed

- Documentation
  - Update README files to list accurate build/deploy steps
  - Doc comments cleanup
- Unable to override CLI flag default for IgnoreDNSErrors from config file
- Linting
  - unhandled error return values
  - file inclusion via variable

## [v0.2.1] - 2020-07-07

### Added

- Enable Dependabot version updates
  - `GitHub Actions`
  - `Go Modules`

### Changed

- Update dependencies
  - `actions/setup-go`
    - `v2.0.3` to `v2.1.0`
  - `actions/setup-node`
    - `v2.0.0` to `v2.1.0`
  - `actions/checkout`
    - `v1` to `v2.3.1`
  - `pelletier/go-toml`
    - `v1.7.0` to `v1.8.0`
  - `apex/log`
    - `v1.1.4` to `v1.4.0`
  - `miekg/dns`
    - `v1.1.29` to `v1.1.30`

## [v0.2.0] - 2020-05-02

### Added

- custom timeout value support
  - new `-timeout`, `-to` flags to supply custom timeout values (in seconds)
  - new `timeout` config file setting

- extend example config file with additional DNS servers
  - add Cloudflare filtered DNS servers
    - 1.1.1.2 (No Malware)
    - 1.1.1.3 (No Malware or Adult Content)
    - see also <https://blog.cloudflare.com/introducing-1-1-1-1-for-families/>

- report round-trip time (aka, response time) in summary output to assist with
  identifying slow DNS servers

### Changed

- effective default query timeout changed from 2s to 10s
  - allow for slower DNS servers to finish their work and respond

## [v0.1.2] - 2020-05-01

### Fixed

- Remove bash shebang from GitHub Actions Workflow files
- README missing requirements, installation instructions

### Changed

- Vendor dependencies

- Update dependencies
  - direct
    - `apex/log` updated to `v1.1.4`
    - `pelletier/go-toml` updated to `v1.7.0`
  - indirect
    - `stretchr/testify`
    - `kr/pretty` replaced with `niemeyer/pretty`
    - `mattn/go-colorable`
    - `gopkg.in/check.v1`

- Linting
  - golangci-lint
    - move settings to external config file
    - enable `scopelint` linter
    - enable `gofmt` linter
    - enable `dogsled` linter
    - switch from build-from-source to binary installation
    - use v1.25.1 release
  - Remove installation step for `gofmt`
  - Remove separate `gofmt` and `golint` calls
    - handled by golangci-lint now

## [v0.1.1] - 2020-03-22

### Added

- Add expanded coverage of config file purpose and supported locations for
  auto-load behavior
- Add simpler flags-only, A record only query

### Fixed

- Fix config file reference to match the filename used in examples

## [v0.1.0] - 2020-03-22

### Added

Features of the initial prototype:

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

Worth noting (in no particular order):

- Command-line flags support via `flag` standard library package
- Go modules (vs classic `GOPATH` setup)
- GitHub Actions Workflows which apply linting and build checks
- Makefile for general use cases (including local linting)
  - Note: See README for available options if building on Windows

[Unreleased]: https://github.com/atc0005/dnsc/compare/v0.6.3...HEAD
[v0.6.3]: https://github.com/atc0005/dnsc/releases/tag/v0.6.3
[v0.6.2]: https://github.com/atc0005/dnsc/releases/tag/v0.6.2
[v0.6.1]: https://github.com/atc0005/dnsc/releases/tag/v0.6.1
[v0.6.0]: https://github.com/atc0005/dnsc/releases/tag/v0.6.0
[v0.5.8]: https://github.com/atc0005/dnsc/releases/tag/v0.5.8
[v0.5.7]: https://github.com/atc0005/dnsc/releases/tag/v0.5.7
[v0.5.6]: https://github.com/atc0005/dnsc/releases/tag/v0.5.6
[v0.5.5]: https://github.com/atc0005/dnsc/releases/tag/v0.5.5
[v0.5.4]: https://github.com/atc0005/dnsc/releases/tag/v0.5.4
[v0.5.3]: https://github.com/atc0005/dnsc/releases/tag/v0.5.3
[v0.5.2]: https://github.com/atc0005/dnsc/releases/tag/v0.5.2
[v0.5.1]: https://github.com/atc0005/dnsc/releases/tag/v0.5.1
[v0.5.0]: https://github.com/atc0005/dnsc/releases/tag/v0.5.0
[v0.4.0]: https://github.com/atc0005/dnsc/releases/tag/v0.4.0
[v0.3.5]: https://github.com/atc0005/dnsc/releases/tag/v0.3.5
[v0.3.4]: https://github.com/atc0005/dnsc/releases/tag/v0.3.4
[v0.3.3]: https://github.com/atc0005/dnsc/releases/tag/v0.3.3
[v0.3.2]: https://github.com/atc0005/dnsc/releases/tag/v0.3.2
[v0.3.1]: https://github.com/atc0005/dnsc/releases/tag/v0.3.1
[v0.3.0]: https://github.com/atc0005/dnsc/releases/tag/v0.3.0
[v0.2.1]: https://github.com/atc0005/dnsc/releases/tag/v0.2.1
[v0.2.0]: https://github.com/atc0005/dnsc/releases/tag/v0.2.0
[v0.1.2]: https://github.com/atc0005/dnsc/releases/tag/v0.1.2
[v0.1.1]: https://github.com/atc0005/dnsc/releases/tag/v0.1.1
[v0.1.0]: https://github.com/atc0005/dnsc/releases/tag/v0.1.0
