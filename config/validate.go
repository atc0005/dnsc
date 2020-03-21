// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/dnsc
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"fmt"
	"strings"

	"github.com/apex/log"
)

// Validate verifies all struct fields have been provided acceptable values
func (c Config) Validate() error {

	// TODO: Ensure this is completely optional
	// if c.configFile == "" {
	// 	return fmt.Errorf("missing fully-qualified path to config file to load")
	// }
	// log.Debugf("c.configFile validates: %#v", c.configFile)

	if c.Servers() == nil || len(c.Servers()) == 0 {
		return fmt.Errorf("one or more DNS servers not provided")
	}
	log.Debugf("c.Servers() validates: (%d entries) %#v", len(c.Servers()), c.Servers())

	if c.cliConfig.IgnoreDNSErrors == nil {
		return fmt.Errorf("c.cliConfig.IgnoreDNSErrors not initialized")
	}
	log.Debug("c.cliConfig.IgnoreDNSErrors initialized")

	if c.fileConfig.IgnoreDNSErrors == nil {
		return fmt.Errorf("c.fileConfig.IgnoreDNSErrors not initialized")
	}
	log.Debug("c.fileConfig.IgnoreDNSErrors initialized")

	if c.Query() == "" {
		return fmt.Errorf("query not provided")
	}
	log.Debugf("c.Query() validates: %#v", c.Query())

	// We'll go ahead and provide a default
	//
	// if c.RequestTypes() == nil {
	// 	return fmt.Errorf("record type not provided")
	// }

	// if not nil, assume that we're dealing with one or more requested record
	// types
	for _, requestType := range c.RequestTypes() {
		switch strings.ToUpper(requestType) {
		case RequestTypeA:
		case RequestTypeAAAA:
		case RequestTypeCNAME:
		case RequestTypeMX:
		default:
			return fmt.Errorf(
				"invalid option %q provided for request type",
				requestType,
			)
		}
	}
	log.Debugf("c.RequestTypes() validates: %#v", c.RequestTypes())

	switch c.LogLevel() {
	case LogLevelFatal:
	case LogLevelError:
	case LogLevelWarn:
	case LogLevelInfo:
	case LogLevelDebug:
	default:
		return fmt.Errorf("invalid option %q provided for log level",
			c.LogLevel())
	}
	log.Debugf("c.LogLevel() validates: %#v", c.LogLevel())

	switch c.LogFormat() {
	case LogFormatCLI:
	case LogFormatJSON:
	case LogFormatLogFmt:
	case LogFormatText:
	case LogFormatDiscard:
	default:
		return fmt.Errorf("invalid option %q provided for log format",
			c.LogFormat())
	}
	log.Debugf("c.LogFormat() validates: %#v", c.LogFormat())

	// Optimist
	log.Debug("All validation checks pass")
	return nil

}
