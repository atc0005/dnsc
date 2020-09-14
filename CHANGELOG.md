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

[Unreleased]: https://github.com/atc0005/dnsc/compare/v0.3.2...HEAD
[v0.3.2]: https://github.com/atc0005/dnsc/releases/tag/v0.3.2
[v0.3.1]: https://github.com/atc0005/dnsc/releases/tag/v0.3.1
[v0.3.0]: https://github.com/atc0005/dnsc/releases/tag/v0.3.0
[v0.2.1]: https://github.com/atc0005/dnsc/releases/tag/v0.2.1
[v0.2.0]: https://github.com/atc0005/dnsc/releases/tag/v0.2.0
[v0.1.2]: https://github.com/atc0005/dnsc/releases/tag/v0.1.2
[v0.1.1]: https://github.com/atc0005/dnsc/releases/tag/v0.1.1
[v0.1.0]: https://github.com/atc0005/dnsc/releases/tag/v0.1.0
