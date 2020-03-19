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
	"github.com/pelletier/go-toml"
)

// Overridden via Makefile for release builds
var version string = "dev build"

// Primarily used with branding
const myAppName string = "dnsc"
const myAppURL string = "https://github.com/atc0005/" + myAppName

const (
	versionFlagHelp = "Whether to display application version and then immediately exit application."
	queryFlagHelp   = "Fully-qualified system to lookup from all provided DNS servers."
)

// Default flag settings if not overridden by user input
const (
	defaultLogLevel string = "info"
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

// Config is a unified set of configuration values for this application. This
// struct is configured via command-line flags or TOML configuration file
// provided by the user.
type Config struct {
	Servers     []string `toml:"dns_servers"`
	Query       string   `toml:"-"`
	ConfigFile  string   `toml:"-"`
	ShowVersion bool     `toml:"-"`
	LogLevel    string   `toml:"log_level"`
}

func (c Config) String() string {
	return fmt.Sprintf(
		"Servers: %v, Query: %q, ConfigFile: %q, ShowVersion: %t, LogLevel: %s",
		c.Servers,
		c.Query,
		c.ConfigFile,
		c.ShowVersion,
		c.LogLevel,
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

	if err := toml.Unmarshal(configFile, c); err != nil {
		return err
	}

	return nil
}

// configureLogging is a wrapper function to enable setting requested logging
// settings.
func (c Config) configureLogging() {
	switch c.LogLevel {
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

	log.SetHandler(cli.New(os.Stdout))
}

// NewConfig is a factory function that produces a new Config object based
// on user provided flag and config file values.
func NewConfig() (*Config, error) {

	config := Config{}

	flag.StringVar(&config.ConfigFile, "config-file", "", "Full path to optional TOML-formatted configuration file. See config.example.toml for a starter template.")

	flag.BoolVar(&config.ShowVersion, "version", false, versionFlagHelp)
	flag.BoolVar(&config.ShowVersion, "v", false, versionFlagHelp+" (shorthand)")

	flag.StringVar(&config.Query, "query", "", queryFlagHelp)
	flag.StringVar(&config.Query, "q", "", queryFlagHelp+" (shorthand)")

	flag.StringVar(
		&config.LogLevel,
		"log-lvl",
		defaultLogLevel,
		"Log message priority filter. Log messages with a lower level are ignored.",
	)

	flag.Usage = flagsUsage()
	flag.Parse()

	// Return immediately if user just wants version details
	if config.ShowVersion {
		return &config, nil
	}

	if err := config.Validate(); err != nil {
		flag.Usage()
		return nil, err
	}

	config.configureLogging()

	// load config file
	log.WithFields(log.Fields{
		"config_file": config.ConfigFile,
	}).Info("attempting to open config file")

	fh, err := os.Open(config.ConfigFile)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	if err := config.LoadConfigFile(fh); err != nil {
		return nil, err
	}

	return &config, nil

}

// Validate verifies all struct fields have been provided acceptable values
func (c Config) Validate() error {

	if c.ConfigFile == "" {
		return fmt.Errorf("missing fully-qualified path to config file to load")
	}

	if c.Servers == nil || len(c.Servers) == 0 {
		return fmt.Errorf("one or more DNS servers not provided")
	}

	if c.Query == "" {
		return fmt.Errorf("query not provided")
	}

	// Optimist
	return nil

}
