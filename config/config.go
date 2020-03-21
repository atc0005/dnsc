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
	"strings"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/apex/log/handlers/discard"
	"github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/logfmt"
	"github.com/apex/log/handlers/text"
	"github.com/miekg/dns"
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
	dnsServerFlagHelp       = "DNS server to submit query against. This flag may be repeated for each additional DNS server to query."
	dnsRequestTypeFlagHelp  = "DNS record type to request when submitting DNS query. This flag may be repeated for each additional DNS record type you wish to request."
)

// Default flag settings if not overridden by user input
const (
	defaultLogLevel              string = "info"
	defaultLogFormat             string = "text"
	defaultDisplayVersionAndExit bool   = false
	defaultIgnoreDNSErrors       bool   = false
	defaultConfigFileName        string = "config.toml"
	defaultRecordType            string = "A"
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

// Supported Request types
const (
	RequestTypeA     string = "A"
	RequestTypeAAAA  string = "AAAA"
	RequestTypeCNAME string = "CNAME"
	RequestTypeMX    string = "MX"
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

// multiValueFlag is a custom type that satisfies the flag.Value interface in
// order to accept multiple values for some of our flags
type multiValueFlag []string

// String returns a comma separated string consisting of all slice elements
func (i *multiValueFlag) String() string {

	// From the `flag` package docs:
	// "The flag package may call the String method with a zero-valued
	// receiver, such as a nil pointer."
	if i == nil {
		return ""
	}

	return strings.Join(*i, ",")
}

// Set is called once by the flag package, in command line order, for each
// flag present
func (i *multiValueFlag) Set(value string) error {
	*i = append(*i, value)
	return nil
}

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
	Servers multiValueFlag `toml:"dns_servers"`

	// RecordTypes is a list of the DNS records that should be requested when
	// submitting queries. This is a a collection of plaintext types provided
	// by the user. See also RecordTypes which is a set of dns.RR types
	// converted from this collection.
	RequestTypes multiValueFlag `toml:"dns_request_types"`

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

func (c Config) String() string {
	return fmt.Sprintf(
		"cliConfig: { Servers: %v, Query: %q, LogLevel: %s, LogFormat: %s, IgnoreDNSErrors: %v, RequestTypes: %v}, "+
			"fileConfig: { Servers: %v, Query: %q, LogLevel: %s, LogFormat: %s, IgnoreDNSErrors: %v, RequestTypes: %v}, "+
			"ConfigFile: %q, ShowVersion: %t,",
		c.cliConfig.Servers,
		c.cliConfig.Query,
		c.cliConfig.LogLevel,
		c.cliConfig.LogFormat,
		c.cliConfig.IgnoreDNSErrors,
		c.cliConfig.RequestTypes,
		c.fileConfig.Servers,
		c.fileConfig.Query,
		c.fileConfig.LogLevel,
		c.fileConfig.LogFormat,
		c.fileConfig.IgnoreDNSErrors,
		c.fileConfig.RequestTypes,
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

// loadConfigFile is a helper function to handle opening a specified config
// file and importing the settings for use
func (c *Config) loadConfigFile(configFile string) (bool, error) {
	// load config file
	log.WithFields(log.Fields{
		"config_file": configFile,
	}).Debug("Attempting to open config file")

	fh, err := os.Open(configFile)
	if err != nil {
		return false, err
	}
	log.Debug("Config file opened")
	defer fh.Close()

	log.Debug("Attempting to import config file")
	result, err := c.ImportConfigFile(fh)
	if err != nil {
		return false, err
	}

	return result, err
}

// localConfigFile returns the potential path to a  local config file or an
// error if one is encountered constructing the path
func (c Config) localConfigFile() (string, error) {

	exePath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("unable to get running executable path to load local config file: %w", err)
	}
	exeDirPath, _ := filepath.Split(exePath)
	localCfgFile := filepath.Join(exeDirPath, defaultConfigFileName)

	log.Debugf("local config file path: %q", localCfgFile)

	// if PathExists(localConfigFile) {
	// 	log.WithFields(log.Fields{
	// 		"local_config_file": localConfigFile,
	// 	}).Info("local config file found")
	// }
	// log.WithFields(log.Fields{
	// 	"local_config_file": localConfigFile,
	// }).Info("local config file not found")

	return localCfgFile, nil
}

// userConfigFile returns the potential path to a  local config file or an
// error if one is encountered constructing the path
func (c Config) userConfigFile() (string, error) {
	// Ubuntu environment:
	// os.UserHomeDir: /home/username
	// os.UserConfigDir: /home/username/.config
	//
	// Windows environment:
	// os.UserHomeDir: C:\Users\username
	// os.UserConfigDir: C:\Users\username\AppData\Roaming
	//
	// Look for:
	// filepath.Join(os.UserConfigDir, "dnsc/config.toml")
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("unable to get user config dir: %w", err)
	}

	userConfigAppDir := filepath.Join(userConfigDir, myAppName)
	userConfigFileFullPath := filepath.Join(userConfigAppDir, defaultConfigFileName)
	log.Debugf("user config file path: %q", userConfigFileFullPath)

	return userConfigFileFullPath, nil
}

// ImportConfigFile reads from an io.Reader and unmarshals a configuration file
// in TOML format into the associated Config struct.
func (c *Config) ImportConfigFile(fh io.Reader) (bool, error) {

	log.Debug("Attempting to read contents of file handle ...")
	configFile, err := ioutil.ReadAll(fh)
	if err != nil {
		return false, err
	}
	log.Debug("Contents of file loaded successfully")

	log.Debug("Attempting to parse TOML file contents")
	// target nested config struct dedicated to TOML config file settings
	if err := toml.Unmarshal(configFile, &c.fileConfig); err != nil {
		return false, err
	}
	log.Debug("Successfully parsed TOML file contents")

	return true, nil
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

// PathExists confirms that the specified path exists
func PathExists(path string) bool {

	// Make sure path isn't empty
	if strings.TrimSpace(path) == "" {
		// DEBUG?
		// WARN?
		// ERROR?
		log.Error("path is empty string")
		return false
	}

	// https://gist.github.com/mattes/d13e273314c3b3ade33f
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		//log.Println("path found")
		return true
	}

	return false

}

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

	//
	// Attempt to load requested config file, fallback to known alternates
	// if user did not specify a config file
	//

	configFiles := make([]string, 0, 3)

	if config.configFile == "" {
		log.Info("User-specified config file not provided")
	} else {
		log.Info("User-specified config file provided, will attempt to load it")
		configFiles = append(configFiles, config.configFile)
	}

	localFile, err := config.localConfigFile()
	if err != nil {
		log.Error("Failed to determine path to local file")
	}
	configFiles = append(configFiles, localFile)

	userConfigFile, err := config.userConfigFile()
	if err != nil {
		log.Error("Failed to determine path to user config file")
	}
	configFiles = append(configFiles, userConfigFile)

	var loadCfgFileSuccess bool
	for _, file := range configFiles {
		log.Infof("Trying to load config file %q", file)
		ok, err := config.loadConfigFile(file)
		if ok {
			loadCfgFileSuccess = true
			// if there were no errors, we're done loading config files
			log.WithFields(log.Fields{
				"config_file": file,
			}).Info("Config file successfully loaded")
			log.Debug("Config file successfully parsed")
			log.Debugf("After loading config file: %v", config.String())
			break
		}

		log.Warnf("Config file %q not found or unable to load", file)
		log.WithFields(log.Fields{
			"error": err,
		}).Debug("")
	}

	if !loadCfgFileSuccess {
		log.Warn("Failed to load config files, relying only on provided flag settings")
	}

	// Apply logging settings based on any provided config file settings
	config.configureLogging()

	log.Debug("Validating configuration ...")
	if err := config.Validate(); err != nil {
		flag.Usage()
		return nil, err
	}
	log.Debug("Configuration validated")

	//log.Debugf("Config object: %v", config.String())

	return &config, nil

}
