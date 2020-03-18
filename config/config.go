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

	"github.com/pelletier/go-toml"
)

// Overridden via Makefile for release builds
var version string = "dev build"

// Primarily used with branding
const myAppName string = "dnsc"
const myAppURL string = "https://github.com/atc0005/dnsc"

const (
	versionFlagHelp = "Whether to display application version and then immediately exit application."
	queryFlagHelp   = "Fully-qualified system to lookup from all provided DNS servers."
)

// Config is a unified set of configuration values for this application. This
// struct is configured via command-line flags or TOML configuration file
// provided by the user.
type Config struct {
	Servers    []string `toml:"servers"`
	Query      string   `toml:"-"`
	ConfigFile string   `toml:"-"`
	Version    bool     `toml:"-"`
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

// NewConfig is a factory function that produces a new Config object based
// on user provided flag and config file values.
func NewConfig() (*Config, error) {

	config := Config{}

	flag.StringVar(&config.ConfigFile, "config-file", "", "Full path to optional TOML-formatted configuration file. See config.example.toml for a starter template.")
	flag.BoolVar(&config.Version, "version", false, versionFlagHelp)
	flag.BoolVar(&config.Version, "v", false, versionFlagHelp+" (shorthand)")

	flag.StringVar(&config.Query, "query", "", queryFlagHelp)
	flag.StringVar(&config.Query, "q", "", queryFlagHelp+" (shorthand)")

	flag.Usage = flagsUsage()
	flag.Parse()

	// Display version info and immediately exit if requested
	if config.Version {
		Branding()
		os.Exit(0)
	}

	// load config file
	fh, err := os.Open(config.ConfigFile)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	if err := config.LoadConfigFile(fh); err != nil {
		return nil, err
	}

	if err := config.Validate(); err != nil {
		flag.Usage()
		return nil, err
	}

	return &config, nil

}

// Validate verifies all struct fields have been provided acceptable values
func (c Config) Validate() error {

	if c.Servers == nil || len(c.Servers) == 0 {
		return fmt.Errorf("one or more DNS servers not provided")
	}

	if c.Query == "" {
		return fmt.Errorf("query not provided")
	}

	if c.ConfigFile == "" {
		return fmt.Errorf("missing fully-qualified path to config file to load")
	}

	// Optimist
	return nil

}
