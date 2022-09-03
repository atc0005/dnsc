// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/dnsc
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

// Run a DNS query concurrently against all servers in a list and provide
// summary of results. This is most useful after moving servers between
// subnets when an IP Address change is expected, but where the change may not
// have propagated between all DNS servers. The summary output is useful for
// spotting systems lagging behind the others.
//
// See our [GitHub repo]:
//
//   - to review documentation (including examples)
//   - for the latest code
//   - to file an issue or submit improvements for review and potential
//     inclusion into the project
//
// [GitHub repo]: https://github.com/atc0005/dnsc
package main
