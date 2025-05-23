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

## [v0.7.15] - 2025-05-16

### Changed

#### Dependency Updates

- (GH-880) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.22.9 to go-ci-oldstable-build-v0.22.10 in /dependabot/docker/builds
- (GH-876) Go Dependency: Bump github.com/miekg/dns from 1.1.65 to 1.1.66
- (GH-867) Go Dependency: Bump golang.org/x/net from 0.39.0 to 0.40.0
- (GH-866) Go Dependency: Bump golang.org/x/sync from 0.13.0 to 0.14.0
- (GH-869) Go Dependency: Bump golang.org/x/sys from 0.32.0 to 0.33.0
- (GH-868) Go Dependency: Bump golang.org/x/tools from 0.32.0 to 0.33.0
- (GH-875) Go Runtime: Bump golang from 1.23.8 to 1.23.9 in /dependabot/docker/go

## [v0.7.14] - 2025-04-14

### Changed

#### Dependency Updates

- (GH-756) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.16 to go-ci-oldstable-build-v0.21.17 in /dependabot/docker/builds
- (GH-846) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.17 to go-ci-oldstable-build-v0.22.9 in /dependabot/docker/builds
- (GH-847) Disable Dependabot automatic PR rebasing
- (GH-781) Go Dependency: Bump github.com/mattn/go-colorable from 0.1.13 to 0.1.14
- (GH-851) Go Dependency: Bump github.com/miekg/dns from 1.1.62 to 1.1.65
- (GH-855) Go Dependency: Bump github.com/pelletier/go-toml/v2 from 2.2.3 to 2.2.4
- (GH-824) Go Dependency: Bump golang.org/x/mod from 0.22.0 to 0.24.0
- (GH-857) Go Dependency: Bump golang.org/x/net from 0.31.0 to 0.39.0
- (GH-849) Go Dependency: Bump golang.org/x/sync from 0.9.0 to 0.13.0
- (GH-850) Go Dependency: Bump golang.org/x/sys from 0.27.0 to 0.32.0
- (GH-856) Go Dependency: Bump golang.org/x/tools from 0.27.0 to 0.32.0
- (GH-844) Go Runtime: Bump golang from 1.22.10 to 1.23.8 in /dependabot/docker/go
- (GH-758) Go Runtime: Bump golang from 1.22.9 to 1.22.10 in /dependabot/docker/go
- (GH-833) go.mod: update minimum Go version to 1.23.0
- (GH-812) Update project to Go 1.23 series

### Fixed

- (GH-861) Fix `copyloopvar` linting errors

## [v0.7.13] - 2024-11-14

### Changed

#### Dependency Updates

- (GH-722) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.12 to go-ci-oldstable-build-v0.21.13 in /dependabot/docker/builds
- (GH-749) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.13 to go-ci-oldstable-build-v0.21.15 in /dependabot/docker/builds
- (GH-752) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.15 to go-ci-oldstable-build-v0.21.16 in /dependabot/docker/builds
- (GH-715) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.9 to go-ci-oldstable-build-v0.21.12 in /dependabot/docker/builds
- (GH-734) Go Dependency: Bump github.com/fatih/color from 1.17.0 to 1.18.0
- (GH-721) Go Dependency: Bump github.com/pelletier/go-toml/v2 from 2.2.2 to 2.2.3
- (GH-709) Go Dependency: Bump golang.org/x/mod from 0.20.0 to 0.21.0
- (GH-739) Go Dependency: Bump golang.org/x/mod from 0.21.0 to 0.22.0
- (GH-712) Go Dependency: Bump golang.org/x/net from 0.28.0 to 0.29.0
- (GH-748) Go Dependency: Bump golang.org/x/net from 0.29.0 to 0.31.0
- (GH-737) Go Dependency: Bump golang.org/x/sync from 0.8.0 to 0.9.0
- (GH-710) Go Dependency: Bump golang.org/x/sys from 0.24.0 to 0.25.0
- (GH-738) Go Dependency: Bump golang.org/x/sys from 0.25.0 to 0.27.0
- (GH-717) Go Dependency: Bump golang.org/x/tools from 0.24.0 to 0.25.0
- (GH-747) Go Dependency: Bump golang.org/x/tools from 0.25.0 to 0.27.0
- (GH-713) Go Runtime: Bump golang from 1.22.6 to 1.22.7 in /dependabot/docker/go
- (GH-736) Go Runtime: Bump golang from 1.22.7 to 1.22.9 in /dependabot/docker/go
- (GH-719) Update project Go version to 1.22.0

## [v0.7.12] - 2024-08-22

### Changed

#### Dependency Updates

- (GH-667) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.4 to go-ci-oldstable-build-v0.21.5 in /dependabot/docker/builds
- (GH-671) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.5 to go-ci-oldstable-build-v0.21.6 in /dependabot/docker/builds
- (GH-673) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.6 to go-ci-oldstable-build-v0.21.7 in /dependabot/docker/builds
- (GH-688) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.7 to go-ci-oldstable-build-v0.21.8 in /dependabot/docker/builds
- (GH-697) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.8 to go-ci-oldstable-build-v0.21.9 in /dependabot/docker/builds
- (GH-693) Go Dependency: Bump github.com/miekg/dns from 1.1.61 to 1.1.62
- (GH-678) Go Dependency: Bump golang.org/x/mod from 0.19.0 to 0.20.0
- (GH-682) Go Dependency: Bump golang.org/x/net from 0.27.0 to 0.28.0
- (GH-676) Go Dependency: Bump golang.org/x/sync from 0.7.0 to 0.8.0
- (GH-677) Go Dependency: Bump golang.org/x/sys from 0.22.0 to 0.23.0
- (GH-690) Go Dependency: Bump golang.org/x/sys from 0.23.0 to 0.24.0
- (GH-683) Go Dependency: Bump golang.org/x/tools from 0.23.0 to 0.24.0
- (GH-699) Go Runtime: Bump golang from 1.21.12 to 1.22.6 in /dependabot/docker/go
- (GH-698) Update project to Go 1.22 series

#### Other

- (GH-674) Push `REPO_VERSION` var into containers for builds

## [v0.7.11] - 2024-07-10

### Changed

#### Dependency Updates

- (GH-641) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.7 to go-ci-oldstable-build-v0.20.8 in /dependabot/docker/builds
- (GH-647) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.8 to go-ci-oldstable-build-v0.21.2 in /dependabot/docker/builds
- (GH-649) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.2 to go-ci-oldstable-build-v0.21.3 in /dependabot/docker/builds
- (GH-653) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.3 to go-ci-oldstable-build-v0.21.4 in /dependabot/docker/builds
- (GH-644) Go Dependency: Bump github.com/miekg/dns from 1.1.59 to 1.1.61
- (GH-655) Go Dependency: Bump golang.org/x/mod from 0.18.0 to 0.19.0
- (GH-661) Go Dependency: Bump golang.org/x/net from 0.26.0 to 0.27.0
- (GH-656) Go Dependency: Bump golang.org/x/sys from 0.21.0 to 0.22.0
- (GH-662) Go Dependency: Bump golang.org/x/tools from 0.22.0 to 0.23.0
- (GH-652) Go Runtime: Bump golang from 1.21.11 to 1.21.12 in /dependabot/docker/go

## [v0.7.10] - 2024-06-07

### Changed

#### Dependency Updates

- (GH-619) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.4 to go-ci-oldstable-build-v0.20.5 in /dependabot/docker/builds
- (GH-621) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.5 to go-ci-oldstable-build-v0.20.6 in /dependabot/docker/builds
- (GH-637) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.6 to go-ci-oldstable-build-v0.20.7 in /dependabot/docker/builds
- (GH-617) Go Dependency: Bump github.com/fatih/color from 1.16.0 to 1.17.0
- (GH-631) Go Dependency: Bump golang.org/x/mod from 0.17.0 to 0.18.0
- (GH-632) Go Dependency: Bump golang.org/x/net from 0.25.0 to 0.26.0
- (GH-633) Go Dependency: Bump golang.org/x/sys from 0.20.0 to 0.21.0
- (GH-634) Go Dependency: Bump golang.org/x/tools from 0.21.0 to 0.22.0
- (GH-635) Go Runtime: Bump golang from 1.21.10 to 1.21.11 in /dependabot/docker/go

### Fixed

- (GH-623) Remove inactive maligned linter
- (GH-624) Fix errcheck linting errors

## [v0.7.9] - 2024-05-13

### Changed

#### Dependency Updates

- (GH-602) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.1 to go-ci-oldstable-build-v0.20.2 in /dependabot/docker/builds
- (GH-610) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.2 to go-ci-oldstable-build-v0.20.3 in /dependabot/docker/builds
- (GH-613) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.3 to go-ci-oldstable-build-v0.20.4 in /dependabot/docker/builds
- (GH-595) Go Dependency: Bump github.com/miekg/dns from 1.1.58 to 1.1.59
- (GH-593) Go Dependency: Bump github.com/pelletier/go-toml/v2 from 2.2.0 to 2.2.1
- (GH-597) Go Dependency: Bump github.com/pelletier/go-toml/v2 from 2.2.1 to 2.2.2
- (GH-606) Go Dependency: Bump golang.org/x/net from 0.24.0 to 0.25.0
- (GH-601) Go Dependency: Bump golang.org/x/sys from 0.19.0 to 0.20.0
- (GH-607) Go Dependency: Bump golang.org/x/tools from 0.20.0 to 0.21.0
- (GH-608) Go Runtime: Bump golang from 1.21.9 to 1.21.10 in /dependabot/docker/go

## [v0.7.8] - 2024-04-11

### Changed

#### Dependency Updates

- (GH-567) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.15.4 to go-ci-oldstable-build-v0.16.0 in /dependabot/docker/builds
- (GH-569) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.16.0 to go-ci-oldstable-build-v0.16.1 in /dependabot/docker/builds
- (GH-570) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.16.1 to go-ci-oldstable-build-v0.19.0 in /dependabot/docker/builds
- (GH-574) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.19.0 to go-ci-oldstable-build-v0.20.0 in /dependabot/docker/builds
- (GH-583) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.0 to go-ci-oldstable-build-v0.20.1 in /dependabot/docker/builds
- (GH-565) Go Dependency: Bump github.com/pelletier/go-toml/v2 from 2.1.1 to 2.2.0
- (GH-581) Go Dependency: Bump golang.org/x/mod from 0.16.0 to 0.17.0
- (GH-580) Go Dependency: Bump golang.org/x/net from 0.22.0 to 0.24.0
- (GH-579) Go Dependency: Bump golang.org/x/sys from 0.18.0 to 0.19.0
- (GH-582) Go Dependency: Bump golang.org/x/tools from 0.19.0 to 0.20.0
- (GH-578) Go Runtime: Bump golang from 1.21.8 to 1.21.9 in /dependabot/docker/go

## [v0.7.7] - 2024-03-08

### Changed

#### Dependency Updates

- (GH-560) Add todo/release label to "Go Runtime" PRs
- (GH-543) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.15.2 to go-ci-oldstable-build-v0.15.3 in /dependabot/docker/builds
- (GH-559) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.15.3 to go-ci-oldstable-build-v0.15.4 in /dependabot/docker/builds
- (GH-540) canary: bump golang from 1.21.6 to 1.21.7 in /dependabot/docker/go
- (GH-537) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.15.0 to go-ci-oldstable-build-v0.15.2 in /dependabot/docker/builds
- (GH-550) Go Dependency: Bump golang.org/x/mod from 0.15.0 to 0.16.0
- (GH-549) Go Dependency: Bump golang.org/x/net from 0.21.0 to 0.22.0
- (GH-548) Go Dependency: Bump golang.org/x/sys from 0.17.0 to 0.18.0
- (GH-551) Go Dependency: Bump golang.org/x/tools from 0.18.0 to 0.19.0
- (GH-556) Go Runtime: Bump golang from 1.21.7 to 1.21.8 in /dependabot/docker/go
- (GH-542) Update Dependabot PR prefixes (redux)
- (GH-541) Update Dependabot PR prefixes
- (GH-539) Update project to Go 1.21 series

### Fixed

- (GH-546) Fix loopvar behavior change linting error

## [v0.7.6] - 2024-02-15

### Changed

#### Dependency Updates

- (GH-524) canary: bump golang from 1.20.13 to 1.20.14 in /dependabot/docker/go
- (GH-503) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.3 to go-ci-oldstable-build-v0.14.4 in /dependabot/docker/builds
- (GH-509) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.4 to go-ci-oldstable-build-v0.14.5 in /dependabot/docker/builds
- (GH-511) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.5 to go-ci-oldstable-build-v0.14.6 in /dependabot/docker/builds
- (GH-525) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.6 to go-ci-oldstable-build-v0.14.9 in /dependabot/docker/builds
- (GH-529) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.9 to go-ci-oldstable-build-v0.15.0 in /dependabot/docker/builds
- (GH-517) go.mod: bump golang.org/x/mod from 0.14.0 to 0.15.0
- (GH-522) go.mod: bump golang.org/x/net from 0.20.0 to 0.21.0
- (GH-523) go.mod: bump golang.org/x/sys from 0.16.0 to 0.17.0
- (GH-531) go.mod: bump golang.org/x/tools from 0.17.0 to 0.18.0

### Fixed

- (GH-533) Fix `unused-parameter` revive linting errors

## [v0.7.5] - 2024-01-19

### Changed

#### Dependency Updates

- (GH-494) canary: bump golang from 1.20.12 to 1.20.13 in /dependabot/docker/go
- (GH-498) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.2 to go-ci-oldstable-build-v0.14.3 in /dependabot/docker/builds
- (GH-488) ghaw: bump github/codeql-action from 2 to 3
- (GH-505) go.mod: bump github.com/miekg/dns from 1.1.57 to 1.1.58
- (GH-485) go.mod: bump github.com/pelletier/go-toml/v2 from 2.1.0 to 2.1.1
- (GH-490) go.mod: bump golang.org/x/sys from 0.15.0 to 0.16.0
- (GH-486) go.mod: bump golang.org/x/tools from 0.16.0 to 0.16.1
- (GH-496) go.mod: bump golang.org/x/tools from 0.16.1 to 0.17.0

## [v0.7.4] - 2023-12-09

### Changed

#### Dependency Updates

- (GH-479) canary: bump golang from 1.20.11 to 1.20.12 in /dependabot/docker/go
- (GH-481) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.1 to go-ci-oldstable-build-v0.14.2 in /dependabot/docker/builds
- (GH-476) go.mod: bump golang.org/x/net from 0.18.0 to 0.19.0
- (GH-477) go.mod: bump golang.org/x/sys from 0.14.0 to 0.15.0
- (GH-475) go.mod: bump golang.org/x/tools from 0.15.0 to 0.16.0

## [v0.7.3] - 2023-11-16

### Changed

#### Dependency Updates

- (GH-455) canary: bump golang from 1.20.10 to 1.20.11 in /dependabot/docker/go
- (GH-413) canary: bump golang from 1.20.7 to 1.20.8 in /dependabot/docker/go
- (GH-436) canary: bump golang from 1.20.8 to 1.20.10 in /dependabot/docker/go
- (GH-442) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.10 to go-ci-oldstable-build-v0.13.12 in /dependabot/docker/builds
- (GH-461) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.12 to go-ci-oldstable-build-v0.14.1 in /dependabot/docker/builds
- (GH-397) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.4 to go-ci-oldstable-build-v0.13.5 in /dependabot/docker/builds
- (GH-400) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.5 to go-ci-oldstable-build-v0.13.6 in /dependabot/docker/builds
- (GH-402) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.6 to go-ci-oldstable-build-v0.13.7 in /dependabot/docker/builds
- (GH-414) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.7 to go-ci-oldstable-build-v0.13.8 in /dependabot/docker/builds
- (GH-422) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.8 to go-ci-oldstable-build-v0.13.9 in /dependabot/docker/builds
- (GH-425) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.9 to go-ci-oldstable-build-v0.13.10 in /dependabot/docker/builds
- (GH-407) ghaw: bump actions/checkout from 3 to 4
- (GH-454) go.mod: bump github.com/fatih/color from 1.15.0 to 1.16.0
- (GH-445) go.mod: bump github.com/mattn/go-isatty from 0.0.19 to 0.0.20
- (GH-417) go.mod: bump github.com/miekg/dns from 1.1.55 to 1.1.56
- (GH-463) go.mod: bump github.com/miekg/dns from 1.1.56 to 1.1.57
- (GH-403) go.mod: bump github.com/pelletier/go-toml/v2 from 2.0.9 to 2.1.0
- (GH-428) go.mod: bump golang.org/x/mod from 0.12.0 to 0.13.0
- (GH-450) go.mod: bump golang.org/x/mod from 0.13.0 to 0.14.0
- (GH-410) go.mod: bump golang.org/x/net from 0.14.0 to 0.15.0
- (GH-438) go.mod: bump golang.org/x/net from 0.15.0 to 0.17.0
- (GH-458) go.mod: bump golang.org/x/net from 0.17.0 to 0.18.0
- (GH-405) go.mod: bump golang.org/x/sys from 0.11.0 to 0.12.0
- (GH-429) go.mod: bump golang.org/x/sys from 0.12.0 to 0.13.0
- (GH-449) go.mod: bump golang.org/x/sys from 0.13.0 to 0.14.0
- (GH-409) go.mod: bump golang.org/x/tools from 0.12.0 to 0.13.0
- (GH-434) go.mod: bump golang.org/x/tools from 0.13.0 to 0.14.0
- (GH-457) go.mod: bump golang.org/x/tools from 0.14.0 to 0.15.0

### Fixed

- (GH-466) Fix goconst linting errors

## [v0.7.2] - 2023-08-17

### Added

- (GH-362) Add initial automated release notes config
- (GH-364) Add initial automated release build workflow

### Changed

- Dependencies
  - `Go`
    - `1.19.11` to `1.20.7`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.11.4` to `go-ci-oldstable-build-v0.13.4`
  - `golang.org/x/net`
    - `v0.12.0` to `v0.13.0`
  - `golang.org/x/tools`
    - `v0.11.0` to `v0.12.0`
- (GH-366) Update Dependabot config to monitor both branches
- (GH-392) Update project to Go 1.20 series

## [v0.7.1] - 2023-07-14

### Overview

- RPM package improvements
- Bug fixes
- Dependency updates
- built using Go 1.19.11
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.7` to `1.19.11`
  - `miekg/dns`
    - `v1.1.53` to `v1.1.55`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.10.3` to `go-ci-oldstable-build-v0.11.4`
  - `pelletier/go-toml`
    - `v2.0.7` to `v2.0.9`
  - `golang.org/x/sys`
    - `v0.6.0` to `v0.10.0`
  - `golang.org/x/net`
    - `v0.8.0` to `v0.12.0`
  - `golang.org/x/mod`
    - `v0.9.0` to `v0.12.0`
  - `golang.org/x/tools`
    - `v0.7.0` to `v0.11.0`
  - `github.com/mattn/go-isatty`
    - `v0.0.18` to `v0.0.19`
- (GH-344) Update vuln analysis GHAW to remove on.push hook
- (GH-347) Restore local CodeQL workflow

### Fixed

- (GH-331) Fix markdownlint MD034 linting error
- (GH-332) Add missing decompression step to README
- (GH-341) Disable depguard linter

## [v0.7.0] - 2023-04-03

### Overview

- Add support for generating DEB, RPM packages
- Build improvements
- Generated binary changes
  - filename patterns
  - compression (~ 66% smaller)
  - executable metadata
- built using Go 1.19.7
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- (GH-303) Generate RPM/DEB packages using nFPM
- (GH-304) Add version details to Windows executables

### Changed

- (GH-302) Switch to semantic versioning (semver) compatible versioning
  pattern
- (GH-305) Makefile: Compress binaries & use static filenames
- (GH-306) Makefile: Refresh recipes to add "standard" set, new
  package-related options
- (GH-307) Build dev/stable releases using go-ci Docker image

## [v0.6.9] - 2023-04-03

### Overview

- Bug fixes
- Build improvements
- GitHub Actions workflows updates
- Dependency updates
- built using Go 1.19.7
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- (GH-286) Add Go Module Validation, Dependency Updates jobs

### Changed

- Dependencies
  - `Go`
    - `1.19.4` to `1.19.7`
  - `miekg/dns`
    - `v1.1.50` to `v1.1.53`
  - `pelletier/go-toml`
    - `v2.0.6` to `v2.0.7`
  - `golang.org/x/sys`
    - `v0.3.0` to `v0.6.0`
  - `golang.org/x/net`
    - `v0.4.0` to `v0.8.0`
  - `golang.org/x/mod`
    - `v0.7.0` to `v0.9.0`
  - `golang.org/x/tools`
    - `v0.4.0` to `v0.7.0`
  - `github.com/mattn/go-isatty`
    - `v0.0.16` to `v0.0.18`
  - `github.com/fatih/color`
    - `v1.13.0` to `v1.15.0`
  - `github.com/go-logfmt/logfmt`
    - `v0.5.1` to `v0.6.0`
- (GH-318) Move config, dqrs packages to internal path
- (GH-296) Drop `Push Validation` workflow
- (GH-297) Rework workflow scheduling
- (GH-299) Remove `Push Validation` workflow status badge

### Fixed

- (GH-315) Update vuln analysis GHAW to use on.push hook

## [v0.6.8] - 2022-12-09

### Overview

- Bugfixes
- Dependency updates
- built using Go 1.19.4
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.1` to `1.19.4`
  - `pelletier/go-toml`
    - `v2.0.5` to `v2.0.6`
  - `github.com/mattn/go-colorable`
    - `v0.1.7` to `v0.1.13`
  - `github.com/mattn/go-isatty`
    - `v0.0.12` to `v0.0.16`
  - `golang.org/x/sys`
    - `v0.0.0-20210630005230-0f9fa26af87c` to `v0.3.0`
  - `golang.org/x/net`
    - `v0.0.0-20210726213435-c6fcb2dbf985` to `v0.4.0`
  - `golang.org/x/mod`
    - `v0.4.2` to `v0.7.0`
  - `golang.org/x/tools`
    - `v0.1.6-0.20210726203631-07bc1bf47fb2` to `v0.4.0`
  - `golang.org/x/xerrors`
    - `v0.0.0-20200804184101-5ec99f83aff1` to
      `v0.0.0-20220907171357-04be3eba64a2`
  - `github.com/fatih/color`
    - `v1.7.0` to `v1.13.0`
  - `github.com/go-logfmt/logfmt`
    - `v0.4.0` to `v0.5.1`
  - `github.com/kr/logfmt`
    - `v0.0.0-20140226030751-b84e30acd515` to
      `v0.0.0-20210122060352-19f9bcb100e6`
  - `github.com/pkg/errors`
    - `v0.8.1` to `v0.9.1`
- (GH-271) Refactor GitHub Actions workflows to import logic

### Fixed

- (GH-278) Fix Makefile Go module base path detection

## [v0.6.7] - 2022-09-20

### Overview

- Bugfixes
- Dependency updates
- built using Go 1.19.1
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.12` to `1.19.1`
  - `pelletier/go-toml`
    - `v2.0.2` to `v2.0.5`
- (GH-262) Update project to Go 1.19
- (GH-265) Update Makefile and GitHub Actions Workflows
- (GH-267) Add CodeQL GitHub Actions Workflow

### Fixed

- (GH-263) Add missing cmd doc file
- (GH-264) Swap io/ioutil package for io package

## [v0.6.6] - 2022-07-21

### Overview

- Add support for additional record type
- Bugfixes
- Dependency updates
- built using Go 1.17.12
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- (GH-250) Add support for NS records

### Changed

- Dependencies
  - `Go`
    - `1.17.10` to `1.17.12`
  - `miekg/dns`
    - `v1.1.49` to `v1.1.50`
  - `pelletier/go-toml`
    - `v2.0.1` to `v2.0.2`

### Fixed

- (GH-254) Update lintinstall Makefile recipe
- (GH-256) Fix Markdownlint references

## [v0.6.5] - 2022-05-20

### Overview

- Bugfixes
- Dependency updates
- built using Go 1.17.10
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.9` to `1.17.10`
  - `miekg/dns`
    - `v1.1.48` to `v1.1.49`
  - `pelletier/go-toml`
    - `v2.0.0` to `v2.0.1`

### Fixed

- (GH-241) Unable to override CLI flag default for Timeout from config file
- (GH-243) `configTemplate.Timeout` field not represented by `Config Stringer`
  implementation
- (GH-245) Expand `Config.Timeout()` debug logging

## [v0.6.4] - 2022-04-28

### Overview

- Dependency updates
- built using Go 1.17.9
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.7` to `1.17.9`
  - `miekg/dns`
    - `v1.1.46` to `v1.1.48`
  - `pelletier/go-toml`
    - `v1.9.4` to `v2.0.0`

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

[Unreleased]: https://github.com/atc0005/dnsc/compare/v0.7.15...HEAD
[v0.7.15]: https://github.com/atc0005/dnsc/releases/tag/v0.7.15
[v0.7.14]: https://github.com/atc0005/dnsc/releases/tag/v0.7.14
[v0.7.13]: https://github.com/atc0005/dnsc/releases/tag/v0.7.13
[v0.7.12]: https://github.com/atc0005/dnsc/releases/tag/v0.7.12
[v0.7.11]: https://github.com/atc0005/dnsc/releases/tag/v0.7.11
[v0.7.10]: https://github.com/atc0005/dnsc/releases/tag/v0.7.10
[v0.7.9]: https://github.com/atc0005/dnsc/releases/tag/v0.7.9
[v0.7.8]: https://github.com/atc0005/dnsc/releases/tag/v0.7.8
[v0.7.7]: https://github.com/atc0005/dnsc/releases/tag/v0.7.7
[v0.7.6]: https://github.com/atc0005/dnsc/releases/tag/v0.7.6
[v0.7.5]: https://github.com/atc0005/dnsc/releases/tag/v0.7.5
[v0.7.4]: https://github.com/atc0005/dnsc/releases/tag/v0.7.4
[v0.7.3]: https://github.com/atc0005/dnsc/releases/tag/v0.7.3
[v0.7.2]: https://github.com/atc0005/dnsc/releases/tag/v0.7.2
[v0.7.1]: https://github.com/atc0005/dnsc/releases/tag/v0.7.1
[v0.7.0]: https://github.com/atc0005/dnsc/releases/tag/v0.7.0
[v0.6.9]: https://github.com/atc0005/dnsc/releases/tag/v0.6.9
[v0.6.8]: https://github.com/atc0005/dnsc/releases/tag/v0.6.8
[v0.6.7]: https://github.com/atc0005/dnsc/releases/tag/v0.6.7
[v0.6.6]: https://github.com/atc0005/dnsc/releases/tag/v0.6.6
[v0.6.5]: https://github.com/atc0005/dnsc/releases/tag/v0.6.5
[v0.6.4]: https://github.com/atc0005/dnsc/releases/tag/v0.6.4
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
