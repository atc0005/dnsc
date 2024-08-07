// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/dnsc
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/apex/log/handlers/discard"
	"github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/logfmt"
	"github.com/apex/log/handlers/text"
)

// Overridden via Makefile for release builds
var version = "dev build"

// ErrVersionRequested indicates that the user requested application version
// information.
var ErrVersionRequested = errors.New("version information requested")

// Primarily used with branding
const myAppName string = "dnsc"
const myAppURL string = "https://github.com/atc0005/" + myAppName

const (
	versionFlagHelp        = "Whether to display application version and then immediately exit application."
	queryFlagHelp          = "Fully-qualified system to lookup from all provided DNS servers."
	logLevelFlagHelp       = "Log message priority filter. Log messages with a lower level are ignored."
	logFormatFlagHelp      = "Log messages are written in this format."
	dnsErrorsFatalFlagHelp = "Whether DNS-related errors should force this application to immediately exit."
	omitTimestampFlagHelp  = "Whether the date/time that results are generated is omitted from the results output."
	configFileFlagHelp     = "Full path to TOML-formatted configuration file. See config.example.toml for a starter template."
	dnsServerFlagHelp      = "DNS server to submit query against. This flag may be repeated for each additional DNS server to query."
	dnsRequestTypeFlagHelp = "DNS query type to use when submitting DNS queries. The default is the 'A' query type. This flag may be repeated for each additional DNS record type you wish to request."
	dnsTimeoutFlagHelp     = "Maximum number of seconds allowed for a DNS query to take before timing out."
	srvProtocolFlagHelp    = "Service Location (SRV) protocols associated with a given domain name as the query string. For example, \"msdcs\" can be specified as the SRV record protocol along with \"example.com\" as the query string to search DNS for \"_ldap._tcp.dc._msdcs.example.com\". This flag may be repeated for each additional SRV protocol that you wish to request records for."
	resultsOutputFlagHelp  = "Specifies whether the results summary output is composed of a single comma-separated line of records for a query, or whether the records are returned one per line."
)

// shorthandFlagSuffix is appended to short flag help text to emphasize that
// the flag is a shorthand version of a longer flag.
const shorthandFlagSuffix = " (shorthand)"

// Default flag settings if not overridden by user input
const (
	defaultLogLevel              string = "info"
	defaultLogFormat             string = "text"
	defaultDisplayVersionAndExit bool   = false
	defaultDNSErrorsFatal        bool   = false
	defaultOmitTimestamp         bool   = false
	defaultConfigFileName        string = "config.toml"
	defaultQueryType             string = "A"
	defaultConfigFile            string = ""
	defaultQuery                 string = ""
	defaultResultsOutput         string = ResultsOutputMultiLine

	// the default timeout is set by the `miekg/dns.dnsTimeout` value, which
	// at the time of this writing is 2 seconds. we override with our own
	// longer default (GH-17), but allow the user to override with their own
	// preference later.
	defaultTimeout int = 10
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
// TODO: Duplicated in dqrs package
const (
	RequestTypeA     string = "A"
	RequestTypeAAAA  string = "AAAA"
	RequestTypeCNAME string = "CNAME"
	RequestTypeMX    string = "MX"
	RequestTypeNS    string = "NS"
	RequestTypePTR   string = "PTR"
	RequestTypeSRV   string = "SRV"
)

// Supported Service Location (SRV) Protocol keywords
const (
	SrvProtocolMSDCS      string = "msdcs"
	SrvProtocolKerberos   string = "kerberos"
	SrvProtocolXMPPServer string = "xmppsrv"
	SrvProtocolXMPPClient string = "xmppclient"
	SrvProtocolSIP        string = "sip"
)

// apex/log Handlers
//
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

// Results summary output display options
// TODO: Duplicated in dqrs package
const (
	ResultsOutputSingleLine string = "single-line"
	ResultsOutputMultiLine  string = "multi-line"
)

// multiValueFlag is a custom type that satisfies the flag.Value interface in
// order to accept multiple values for some of our flags
type multiValueFlag []string

// String returns a comma separated string consisting of all slice elements
func (mvf *multiValueFlag) String() string {

	// From the `flag` package docs:
	// "The flag package may call the String method with a zero-valued
	// receiver, such as a nil pointer."
	if mvf == nil {
		return ""
	}

	return strings.Join(*mvf, ",")
}

// Set is called once by the flag package, in command line order, for each
// flag present
func (mvf *multiValueFlag) Set(value string) error {

	// for _, flagVal := range *mvf {
	// 	if flagVal == value {
	// 		return nil
	// 	}
	// }
	// *mvf = append(*mvf, value)
	for i := range *mvf {
		// dereference the pointer slice in order to index into the slice
		// https://stackoverflow.com/questions/38468258/why-is-indexing-on-the-slice-pointer-not-allowed-in-golang
		// https://stackoverflow.com/questions/28709254/how-to-access-elements-from-slice-using-index-which-is-passed-by-reference-in-go
		if (*mvf)[i] == value {
			return nil
		}
	}
	*mvf = append(*mvf, value)

	return nil
}

// Config is a unified set of configuration values for this application. This
// struct is configured via command-line flags or TOML configuration file
// provided by the user. The majority of values held by this object are
// intended to be retrieved via "Getter" methods.
type Config struct {

	// configFile represents the fully-qualified path to a configuration file
	// consulted for settings not provided via CLI flags
	configFile string `toml:"-"`

	// Use our template, define distinct collections of configuration settings
	cliConfig  configTemplate
	fileConfig configTemplate

	// showVersion is a flag indicating whether the user opted to display only
	// the version string and then immediately exit the application
	showVersion bool `toml:"-"`
}

// configTemplate is our base configuration template used to collect values
// specified by various configuration sources
type configTemplate struct {

	// LogLevel is the chosen logging level
	LogLevel string `toml:"log_level"`

	// LogFormat controls which output format is used for log messages
	// generated by this application. This value is from a smaller subset
	// of the formats supported by the third-party leveled-logging package
	// used by this application.
	LogFormat string `toml:"log_format"`

	// ResultsOutput specifies whether the results summary is composed of a
	// single comma-separated line of records for a query, or whether the
	// records are returned one per line.
	ResultsOutput string `toml:"results_output"`

	// Query represents the FQDN query strings submitted to each DNS server
	Query string `toml:"query"`

	// Servers is a list of the DNS servers used by this application. Most
	// commonly set in a configuration file due to the number of servers used
	// for testing queries.
	Servers multiValueFlag `toml:"dns_servers"`

	// QueryTypes is a list of the DNS records that will be requested when
	// submitting queries. This is a a collection of keywords provided by the
	// user. Each keyword/string corresponds to an internal Resource Record
	// type (dns.RR).
	QueryTypes multiValueFlag `toml:"dns_query_types"`

	// SrvProtocols is a list of the Service Location (SRV) record protocol
	// keywords associated with a given domain name.
	SrvProtocols multiValueFlag `toml:"dns_srv_protocols"`

	// Timeout is the number of seconds allowed for a DNS query to complete
	// before it times out.
	Timeout int `toml:"timeout"`

	// DNSErrorsFatal is a boolean *pointer* flag indicating whether
	// individual DNS errors should be ignored. If enabled, this setting
	// allows query-related DNS errors with one host to not block queries
	// against remaining DNS servers. This can be useful to work around
	// failures with one server in a pool of many.
	DNSErrorsFatal bool `toml:"dns_errors_fatal"`

	// OmitTimestamp specifies whether the date & time for when the output is
	// generated is omitted from the results.
	OmitTimestamp bool `toml:"omit_timestamp"`
}

func (c Config) String() string {
	return fmt.Sprintf(
		"cliConfig: { Servers: %v, Query: %q, LogLevel: %s, LogFormat: %s, "+
			"ResultsOutput: %s, DNSErrorsFatal: %v, OmitTimestamp: %v, "+
			"QueryTypes: %v, SrvProtocols: %v, Timeout: %v}, "+
			"fileConfig: { Servers: %v, Query: %q, LogLevel: %s, "+
			"LogFormat: %s, ResultsOutput: %s, DNSErrorsFatal: %v, "+
			"OmitTimestamp: %v, QueryTypes: %v, SrvProtocols: %v, "+
			"Timeout: %v}, "+
			"ConfigFile: %q, ShowVersion: %t,",
		c.cliConfig.Servers,
		c.cliConfig.Query,
		c.cliConfig.LogLevel,
		c.cliConfig.LogFormat,
		c.cliConfig.ResultsOutput,
		c.cliConfig.DNSErrorsFatal,
		c.cliConfig.OmitTimestamp,
		c.cliConfig.QueryTypes,
		c.cliConfig.SrvProtocols,
		c.cliConfig.Timeout,
		c.fileConfig.Servers,
		c.fileConfig.Query,
		c.fileConfig.LogLevel,
		c.fileConfig.LogFormat,
		c.fileConfig.ResultsOutput,
		c.fileConfig.DNSErrorsFatal,
		c.fileConfig.OmitTimestamp,
		c.fileConfig.QueryTypes,
		c.fileConfig.SrvProtocols,
		c.fileConfig.Timeout,
		c.configFile,
		c.showVersion,
	)
}

// Branding is responsible for emitting application name, version and origin
func Branding() {
	_, _ = fmt.Fprintf(flag.CommandLine.Output(), "\n%s %s\n%s\n\n", myAppName, version, myAppURL)
}

// flagsUsage displays branding information and general usage details
func flagsUsage() func() {

	return func() {

		myBinaryName := filepath.Base(os.Args[0])

		Branding()

		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage of \"%s\":\n",
			myBinaryName,
		)
		flag.PrintDefaults()

	}
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
		// log.Println("path found")
		return true
	}

	return false

}

// SrvProtocolTmplLookup looks up the protocol record template associated with
// a given protocol keyword. An error is returned if support is not provided
// for a specific protocol keyword.
func SrvProtocolTmplLookup(keyword string) (string, error) {

	if strings.TrimSpace(keyword) == "" {
		return "", fmt.Errorf("missing SRV protocol keyword")

	}

	// valid keywords/protocol strings
	srvProtocolIdx := map[string]string{
		SrvProtocolMSDCS:      "_ldap._tcp.dc._msdcs.%s",
		SrvProtocolKerberos:   "_kerberos._tcp.%s",
		SrvProtocolXMPPServer: "_xmpp-server._tcp.%s",
		SrvProtocolXMPPClient: "_xmpp-client._tcp.%s",
		SrvProtocolSIP:        "_sip._tcp.%s",
	}

	val, ok := srvProtocolIdx[keyword]
	if !ok {
		return "", fmt.Errorf(
			"unsupported SRV protocol keyword specified: %s. ",
			keyword,
		)
	}

	return val, nil

}

// NewConfig is a factory function that produces a new Config object based
// on user provided flag and config file values.
func NewConfig() (*Config, error) {

	var config Config

	config.handleFlagsConfig()

	// Apply initial logging settings based on any provided CLI flags
	config.configureLogging()

	log.Debugf("After parsing flags: %v", config.String())

	// Return immediately if user just wants version details
	if config.ShowVersion() {
		return &config, ErrVersionRequested
	}

	//
	// Attempt to load requested config file first. If not provided or if
	// unable to use, attempt to load config file from known alternative
	// locations.
	//

	configFiles := make([]string, 0, 3)

	if config.configFile == "" {
		log.Debug("User-specified config file not provided")
	} else {
		log.Debug("User-specified config file provided, will attempt to load it")
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
		log.Debugf("Trying to load config file %q", file)
		ok, err := config.loadConfigFile(file)
		if ok {
			loadCfgFileSuccess = true
			// if there were no errors, we're done loading config files
			log.WithFields(log.Fields{
				"config_file": file,
			}).Debug("Config file successfully loaded")
			log.Debug("Config file successfully parsed")
			log.Debugf("After loading config file: %v", config.String())
			break
		}

		log.WithFields(log.Fields{
			"error": err,
			"file":  file,
		}).Debug("Config file not found or unable to load")
	}

	if !loadCfgFileSuccess {
		log.Debug("Failed to load config files, relying only on provided flag settings")
	}

	// Apply logging settings based on any provided config file settings
	config.configureLogging()

	// If SRV record protocol(s) specified, assume that user has specified or
	// wishes to also specify SRV record type; add to the list of query types
	// as a convenience to the user.
	if len(config.SrvProtocols()) > 0 {

		switch {
		case config.cliConfig.QueryTypes != nil:
			err := config.cliConfig.QueryTypes.Set(strings.ToLower(RequestTypeSRV))
			if err != nil {
				return nil, fmt.Errorf(
					"failed to assert SRV record query type for specified SRV protocols: %w",
					err,
				)
			}
		case config.fileConfig.QueryTypes != nil:
			err := config.fileConfig.QueryTypes.Set(strings.ToLower(RequestTypeSRV))
			if err != nil {
				return nil, fmt.Errorf(
					"failed to assert SRV record query type for specified SRV protocols: %w",
					err,
				)
			}
		}

	}

	log.Debug("Validating configuration ...")
	if err := config.Validate(); err != nil {
		flag.Usage()
		return nil, err
	}
	log.Debug("Configuration validated")

	// log.Debugf("Config object: %v", config.String())

	return &config, nil

}
