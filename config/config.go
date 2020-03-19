// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/dnsc
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

// Package config provides types and functions to collect, validate and apply
// user-provided settings.
package config

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/apex/log/handlers/discard"
	"github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/logfmt"
	"github.com/apex/log/handlers/text"
	"github.com/pelletier/go-toml"
)

// Overridden via Makefile for release builds
var version string = "dev build"

// Primarily used with branding
const myAppName string = "dnsc"
const myAppURL string = "https://github.com/atc0005/" + myAppName

const (
	versionFlagHelp         = "Whether to display application version and then immediately exit application."
	queryFlagHelp           = "Fully-qualified system to lookup from all provided DNS servers."
	logLevelFlagHelp        = "Log message priority filter. Log messages with a lower level are ignored."
	logFormatFlagHelp       = "Log messages are written in this format"
	ignoreDNSErrorsFlagHelp = "Whether DNS-related errors with one server should be ignored in order to try other DNS servers in the list."
	configFileFlagHelp      = "Full path to TOML-formatted configuration file. See config.example.toml for a starter template."
)

// Default flag settings if not overridden by user input
const (
	defaultLogLevel              string = "info"
	defaultLogFormat             string = "text"
	defaultDisplayVersionAndExit bool   = false
	defaultIgnoreDNSErrors       bool   = false
)

// Log levels
const (
	// https://godoc.org/github.com/apex/log#Level

	// LogLevelFatal is used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	LogLevelFatal string = "fatal"

	// LogLevelError is for errors that should definitely be noted.
	LogLevelError string = "error"

	// LogLevelWarn is for non-critical entries that deserve eyes.
	LogLevelWarn string = "warn"

	// LogLevelInfo is for general application operational entries.
	LogLevelInfo string = "info"

	// LogLevelDebug is for debug-level messages and is usually enabled
	// when debugging. Very verbose logging.
	LogLevelDebug string = "debug"
)

// 	apex/log Handlers
// ---------------------------------------------------------
// cli - human-friendly CLI output
// discard - discards all logs
// es - Elasticsearch handler
// graylog - Graylog handler
// json - JSON output handler
// kinesis - AWS Kinesis handler
// level - level filter handler
// logfmt - logfmt plain-text formatter
// memory - in-memory handler for tests
// multi - fan-out to multiple handlers
// papertrail - Papertrail handler
// text - human-friendly colored output
// delta - outputs the delta between log calls and spinner
const (
	// LogFormatCLI provides human-friendly CLI output
	LogFormatCLI string = "cli"

	// LogFormatJSON provides JSON output
	LogFormatJSON string = "json"

	// LogFormatLogFmt provides logfmt plain-text output
	LogFormatLogFmt string = "logfmt"

	// LogFormatText provides human-friendly colored output
	LogFormatText string = "text"

	// LogFormatDiscard discards all logs
	LogFormatDiscard string = "discard"
)

// Config is a unified set of configuration values for this application. This
// struct is configured via command-line flags or TOML configuration file
// provided by the user. The majority of values held by this object are
// intended to be retrieved via "Getter" methods.
type Config struct {

	// Use our template, define distinct collections of configuration settings
	cliConfig  configTemplate
	fileConfig configTemplate

	// configFile represents the fully-qualified path to a configuration file
	// consulted for settings not provided via CLI flags
	configFile string `toml:"-"`

	// showVersion is a flag indicating whether the user opted to display only
	// the version string and then immediately exit the application
	showVersion bool `toml:"-"`
}

// configTemplate is our base configuration template used to collect values
// specified by various configuration sources
type configTemplate struct {

	// IgnoreDNSErrors is a boolean *pointer* flag indicating whether
	// individual DNS errors should be ignored. If enabled, this setting
	// allows query-related DNS errors with one host to not block queries
	// against remaining DNS servers. This can be useful to work around
	// failures with one server in a pool of many.
	IgnoreDNSErrors *bool `toml:"ignore_dns_errors"`

	// Servers is a list of the DNS servers used by this application. Most
	// commonly set in a configuration file due to the number of servers used
	// for testing queries.
	Servers []string `toml:"dns_servers"`

	// Query represents the FQDN query strings submitted to each DNS server
	Query string `toml:"query"`

	// LogLevel is the chosen logging level
	LogLevel string `toml:"log_level"`

	// LogFormat controls which output format is used for log messages
	// generated by this application. This value is from a smaller subset
	// of the formats supported by the third-party leveled-logging package
	// used by this application.
	LogFormat string `toml:"log_format"`
}

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

func (c Config) String() string {
	return fmt.Sprintf(
		"cliConfig: { Servers: %v, Query: %q, LogLevel: %s, LogFormat: %s, IgnoreDNSErrors: %v}, "+
			"fileConfig: { Servers: %v, Query: %q, LogLevel: %s, LogFormat: %s, IgnoreDNSErrors: %v}, "+
			"ConfigFile: %q, ShowVersion: %t,",
		c.cliConfig.Servers,
		c.cliConfig.Query,
		c.cliConfig.LogLevel,
		c.cliConfig.LogFormat,
		c.cliConfig.IgnoreDNSErrors,
		c.fileConfig.Servers,
		c.fileConfig.Query,
		c.fileConfig.LogLevel,
		c.fileConfig.LogFormat,
		c.fileConfig.IgnoreDNSErrors,
		c.configFile,
		c.showVersion,
	)
}

// Branding is responsible for emitting application name, version and origin
func Branding() {
	fmt.Fprintf(flag.CommandLine.Output(), "\n%s %s\n%s\n\n", myAppName, version, myAppURL)
}

// flagsUsage displays branding information and general usage details
func flagsUsage() func() {

	return func() {

		myBinaryName := filepath.Base(os.Args[0])

		Branding()

		fmt.Fprintf(flag.CommandLine.Output(), "Usage of \"%s\":\n",
			myBinaryName,
		)
		flag.PrintDefaults()

	}
}

// LoadConfigFile reads from an io.Reader and unmarshals a configuration file
// in TOML format into the associated Config struct.
func (c *Config) LoadConfigFile(fh io.Reader) error {

	configFile, err := ioutil.ReadAll(fh)
	if err != nil {
		return err
	}

	// target nested config struct dedicated to TOML config file settings
	if err := toml.Unmarshal(configFile, &c.fileConfig); err != nil {
		return err
	}

	return nil
}

// configureLogging is a wrapper function to enable setting requested logging
// settings.
func (c Config) configureLogging() {

	switch c.LogLevel() {
	case LogLevelFatal:
		log.SetLevel(log.FatalLevel)
	case LogLevelError:
		log.SetLevel(log.ErrorLevel)
	case LogLevelWarn:
		log.SetLevel(log.WarnLevel)
	case LogLevelInfo:
		log.SetLevel(log.InfoLevel)
	case LogLevelDebug:
		log.SetLevel(log.DebugLevel)
	}

	switch c.LogFormat() {
	case LogFormatText:
		log.SetHandler(text.New(os.Stdout))
	case LogFormatCLI:
		log.SetHandler(cli.New(os.Stdout))
	case LogFormatLogFmt:
		log.SetHandler(logfmt.New(os.Stdout))
	case LogFormatJSON:
		log.SetHandler(json.New(os.Stdout))
	case LogFormatDiscard:
		log.SetHandler(discard.New())
	}

}

// NewBaseConfig returns a bare-minimum initialized config object for further
// customization before returning to a caller
func NewBaseConfig() Config {

	// we have to explicitly initialize our IgnoreDNSErrors pointer field
	// to prevent "invalid memory address or nil pointer deference" panics
	return Config{
		cliConfig: configTemplate{
			IgnoreDNSErrors: new(bool),
		},
		fileConfig: configTemplate{
			IgnoreDNSErrors: new(bool),
		},
	}
}

// handleFlagsConfig wraps flag setup code into a bundle for potential ease of
// use and future testability
func (c *Config) handleFlagsConfig() {

	log.Debugf("Before parsing flags: %v", c.String())

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

// NewConfig is a factory function that produces a new Config object based
// on user provided flag and config file values.
func NewConfig() (*Config, error) {

	config := NewBaseConfig()

	config.handleFlagsConfig()

	// Apply initial logging settings based on any provided CLI flags
	config.configureLogging()

	log.Debugf("After parsing flags: %v", config.String())

	// Return immediately if user just wants version details
	if config.ShowVersion() {
		return &config, nil
	}

	// load config file
	log.WithFields(log.Fields{
		"config_file": config.configFile,
	}).Debug("Attempting to open config file")

	fh, err := os.Open(config.configFile)
	if err != nil {
		return nil, err
	}
	log.Debug("Config file opened")
	defer fh.Close()

	if err := config.LoadConfigFile(fh); err != nil {
		return nil, err
	}

	// Apply logging settings based on any provided config file settings
	config.configureLogging()

	log.Debug("Config file successfully parsed")

	log.Debugf("After loading config file: %v", config.String())

	log.Debug("Validating configuration after importing config file")
	if err := config.Validate(); err != nil {
		flag.Usage()
		return nil, err
	}
	log.Debug("Configuration validated")

	//log.Debugf("Config object: %v", config.String())

	return &config, nil

}

// Validate verifies all struct fields have been provided acceptable values
func (c Config) Validate() error {

	if c.configFile == "" {
		return fmt.Errorf("missing fully-qualified path to config file to load")
	}
	log.Debugf("c.configFile validates: %#v", c.configFile)

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
