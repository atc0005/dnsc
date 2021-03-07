// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/dnsc
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"time"

	"github.com/apex/log"
)

// Servers returns a slice of configured DNS server entries or nil if DNS
// server entries were not provided. CLI flag values take precedence if
// provided.
func (c Config) Servers() []string {

	switch {
	case c.cliConfig.Servers != nil:
		return c.cliConfig.Servers
	case c.fileConfig.Servers != nil:
		return c.fileConfig.Servers
	default:
		return nil
	}
}

// Query returns the user-provided DNS server query or the default value if
// DNS server query was not provided. CLI flag values take precedence if
// provided.
func (c Config) Query() string {

	switch {
	case c.cliConfig.Query != "":
		return c.cliConfig.Query
	case c.fileConfig.Query != "":
		return c.fileConfig.Query
	default:
		return defaultQuery
	}
}

// LogLevel returns the user-provided logging level or the default value if
// not provided. CLI flag values take precedence if provided.
func (c Config) LogLevel() string {

	switch {
	case c.cliConfig.LogLevel != "":
		return c.cliConfig.LogLevel
	case c.fileConfig.LogLevel != "":
		return c.fileConfig.LogLevel
	default:
		return defaultLogLevel
	}
}

// LogFormat returns the user-provided logging format or empty string if not
// provided. CLI flag values take precedence if provided.
func (c Config) LogFormat() string {

	switch {
	case c.cliConfig.LogFormat != "":
		return c.cliConfig.LogFormat
	case c.fileConfig.LogFormat != "":
		return c.fileConfig.LogFormat
	default:
		return defaultLogFormat
	}
}

// ShowVersion returns the user-provided choice of displaying the application
// version and exiting or the default value for this choice.
func (c Config) ShowVersion() bool {
	return c.showVersion
}

// QueryTypes returns the user-provided choice of which DNS record types to
// request when submitting queries. If not set, defaults to A record type.
func (c Config) QueryTypes() []string {

	switch {
	case c.cliConfig.QueryTypes != nil:
		return c.cliConfig.QueryTypes
	case c.fileConfig.QueryTypes != nil:
		return c.fileConfig.QueryTypes
	default:
		log.Debugf("Requested record types not specified, using default: %q",
			defaultQueryType)
		return []string{defaultQueryType}
	}
}

// SrvProtocols returns the user-provided choice of which Service Location
// (SRV) record protocols should be used with a given query string. This is a
// "shortcut" of sorts that allows specifying one or more short strings
// instead of providing a much longer, unwieldy formatted string as the query
// value as separate queries.
//
// For example, "msdcs" can be specified as a SRV record protocol along with
// "example.com" as the query string to search DNS for
// "_ldap._tcp.dc._msdcs.example.com".
func (c Config) SrvProtocols() []string {
	switch {
	case c.cliConfig.SrvProtocols != nil:
		return c.cliConfig.SrvProtocols
	case c.fileConfig.SrvProtocols != nil:
		return c.fileConfig.SrvProtocols
	default:
		log.Debug("Requested SRV protocol not specified, skipping SRV protocol templating")

		return nil
	}
}

// Timeout returns the user-provided choice of what timeout value to use for
// DNS queries. If not set, returns the default value for our application.
func (c Config) Timeout() time.Duration {

	// FIXME: Initial implementation for GH-17; will need to be revisited
	// alongside the work on GH-10.
	switch {
	case c.cliConfig.Timeout != 0:
		return time.Duration(c.cliConfig.Timeout) * time.Second
	case c.fileConfig.Timeout != 0:
		return time.Duration(c.fileConfig.Timeout) * time.Second
	default:
		log.Debugf("Requested timeout value not specified, using default: %v",
			defaultTimeout)
		return time.Duration(defaultTimeout) * time.Second
	}
}

// DNSErrorsFatal returns the user-provided choice of whether DNS-related
// errors should be fatal or the default value for this choice.
func (c Config) DNSErrorsFatal() bool {
	switch {
	case c.cliConfig.DNSErrorsFatal:
		return c.cliConfig.DNSErrorsFatal
	case c.fileConfig.DNSErrorsFatal:
		return c.fileConfig.DNSErrorsFatal
	default:
		return defaultDNSErrorsFatal
	}
}
