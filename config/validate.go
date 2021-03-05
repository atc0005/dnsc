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

	if c.Query() == "" {
		return fmt.Errorf("query not provided")
	}
	log.Debugf("c.Query() validates: %#v", c.Query())

	// We'll go ahead and provide a default
	//
	// if c.QueryTypes() == nil {
	// 	return fmt.Errorf("record type not provided")
	// }

	// if not nil, assume that we're dealing with one or more requested record
	// types
	for _, queryType := range c.QueryTypes() {
		switch strings.ToUpper(queryType) {
		case RequestTypeA:
		case RequestTypeAAAA:
		case RequestTypeCNAME:
		case RequestTypeMX:
		case RequestTypePTR:
		case RequestTypeSRV:
		default:
			return fmt.Errorf(
				"invalid option %q provided for request type",
				queryType,
			)
		}
	}
	log.Debugf("c.QueryTypes() validates: %#v", c.QueryTypes())

	switch {
	case len(c.SrvProtocols()) > 0:

		// Perfectly acceptable to not specify a SRV protocol, but if specified,
		// limit provided keywords to a valid list.
		for _, queryType := range c.QueryTypes() {
			if strings.ToUpper(queryType) == RequestTypeSRV {
				for _, srvProtocol := range c.SrvProtocols() {
					_, err := SrvProtocolTmplLookup(srvProtocol)
					if err != nil {
						return err
					}
				}
			}
		}
		log.Debugf("c.SrvProtocols() validation; valid keywords provided: %#v", c.SrvProtocols())

		// The SRV record query type should have been included earlier as part
		// of creating a new configuration object if the user specifies a SRV
		// protocol. This action is performed as a convenience to the user. We
		// assert that this is true here as the query type is needed for later
		// logic checks.
		var srvTypeSpecified bool
		for _, queryType := range c.QueryTypes() {
			if strings.ToUpper(queryType) == RequestTypeSRV {
				srvTypeSpecified = true
				for _, srvProtocol := range c.SrvProtocols() {
					_, err := SrvProtocolTmplLookup(srvProtocol)
					if err != nil {
						return err
					}
				}
			}
		}

		if !srvTypeSpecified {
			return fmt.Errorf("SRV record protocol specified, but SRV record type NOT specified")
		}
		log.Debugf("c.SrvProtocols() validation; SRV query type present: %#v", c.QueryTypes())

		log.Debug("c.SrvProtocols() validates: all checks pass")
	default:

		log.Debug("c.SrvProtocols() validates: No SRV record protocols specified")
	}

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
