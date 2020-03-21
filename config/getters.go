// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/dnsc
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"strings"

	"github.com/apex/log"
	"github.com/miekg/dns"
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

// Query returns the user-provided DNS server query or empty string if DNS
// server query was not provided. CLI flag values take precedence if provided.
func (c Config) Query() string {

	switch {
	case c.cliConfig.Query != "":
		return c.cliConfig.Query
	case c.fileConfig.Query != "":
		return c.fileConfig.Query
	default:
		return ""
	}
}

// LogLevel returns the user-provided logging level or empty string if not
// provided. CLI flag values take precedence if provided.
func (c Config) LogLevel() string {

	switch {
	case c.cliConfig.LogLevel != "":
		return c.cliConfig.LogLevel
	case c.fileConfig.LogLevel != "":
		return c.fileConfig.LogLevel
	default:
		return ""
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
		return ""
	}
}

// ShowVersion returns the user-provided choice of displaying the application
// version and exiting or the default value for this choice.
func (c Config) ShowVersion() bool {
	return c.showVersion
}

// RequestTypes returns the user-provided choice of which DNS record types to
// request when submitting queries. If not set, defaults to A record type.
func (c Config) RequestTypes() []string {

	switch {
	case c.cliConfig.RequestTypes != nil:
		return c.cliConfig.RequestTypes
	case c.fileConfig.RequestTypes != nil:
		return c.fileConfig.RequestTypes
	default:
		log.Debugf("Requested record types not specified, using default: %q",
			defaultRecordType)
		return []string{defaultRecordType}
	}
}

// RecordTypes converts the requested string keys which represent
// user-requested DNS record types into the native dns package types needed to
// submit queries for those record types.
func (c Config) RecordTypes() []uint16 {

	requestedTypes := c.RequestTypes()
	if requestedTypes == nil || len(requestedTypes) < 1 {
		return nil
	}

	// we have at least one record type to convert
	recordTypes := make([]uint16, 0, len(requestedTypes))

	// at this point we can be confident that we are only working with strings
	// that fit within our supported collection *and* which properly map to
	// dns.RR types
	for _, requestedType := range requestedTypes {
		if recordType, ok := dns.StringToType[strings.ToUpper(requestedType)]; ok {
			recordTypes = append(recordTypes, recordType)
		}
	}

	return recordTypes

}

// IgnoreDNSErrors returns the user-provided choice of ignoring DNS-related
// errors or the default value for this choice.
func (c Config) IgnoreDNSErrors() bool {
	switch {

	// if not nil, a choice was made; use set choice, otherwise return default
	case c.cliConfig.IgnoreDNSErrors != nil:
		return *c.cliConfig.IgnoreDNSErrors
	case c.fileConfig.IgnoreDNSErrors != nil:
		return *c.fileConfig.IgnoreDNSErrors
	default:
		return defaultIgnoreDNSErrors
	}
}
