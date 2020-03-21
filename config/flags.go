// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/dnsc
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"flag"

	"github.com/apex/log"
)

// handleFlagsConfig wraps flag setup code into a bundle for potential ease of
// use and future testability
func (c *Config) handleFlagsConfig() {

	log.Debugf("Before parsing flags: %v", c.String())

	flag.Var(&c.cliConfig.Servers, "ds", dnsServerFlagHelp+" (shorthand)")
	flag.Var(&c.cliConfig.Servers, "dns-server", dnsServerFlagHelp)

	flag.Var(&c.cliConfig.RequestTypes, "t", dnsRequestTypeFlagHelp+" (shorthand)")
	flag.Var(&c.cliConfig.RequestTypes, "type", dnsRequestTypeFlagHelp)

	flag.StringVar(&c.configFile, "cf", "", configFileFlagHelp+" (shorthand)")
	flag.StringVar(&c.configFile, "config-file", "", configFileFlagHelp)

	flag.BoolVar(&c.showVersion, "version", defaultDisplayVersionAndExit, versionFlagHelp)
	flag.BoolVar(&c.showVersion, "v", defaultDisplayVersionAndExit, versionFlagHelp+" (shorthand)")

	flag.BoolVar(c.cliConfig.IgnoreDNSErrors, "ignore-dns-errors", defaultIgnoreDNSErrors, ignoreDNSErrorsFlagHelp)
	flag.BoolVar(c.cliConfig.IgnoreDNSErrors, "ide", defaultIgnoreDNSErrors, ignoreDNSErrorsFlagHelp+" (shorthand)")

	flag.StringVar(&c.cliConfig.Query, "query", "", queryFlagHelp)
	flag.StringVar(&c.cliConfig.Query, "q", "", queryFlagHelp+" (shorthand)")

	// create shorter and longer logging level flag options
	flag.StringVar(&c.cliConfig.LogLevel, "ll", defaultLogLevel, logLevelFlagHelp)
	flag.StringVar(&c.cliConfig.LogLevel, "log-level", defaultLogLevel, logLevelFlagHelp)

	// create shorter and longer logging format flag options
	flag.StringVar(&c.cliConfig.LogFormat, "lf", defaultLogFormat, logFormatFlagHelp)
	flag.StringVar(&c.cliConfig.LogFormat, "log-format", defaultLogFormat, logFormatFlagHelp)

	flag.Usage = flagsUsage()
	flag.Parse()
}
