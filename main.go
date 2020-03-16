// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/dnsc
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/miekg/dns"
	"github.com/pelletier/go-toml"
)

// Overridden via Makefile for release builds
var version string = "dev build"

// Primarily used with branding
const myAppName string = "dnsc"
const myAppURL string = "https://github.com/atc0005/dnsc"

// Config is a unified set of configuration values for this application. This
// struct is configured via command-line flags or TOML configuration file
// provided by the user.
type Config struct {
	DNSServers     []string `toml:"dns_servers"`
	Query          string   `toml:"-"`
	ExpectedAnswer string   `toml:"-"`
	ConfigFile     string   `toml:"-"`
	Version        bool     `toml:"-"`
}

func main() {

	const (
		versionFlagHelp        = "Whether to display application version and then immediately exit application."
		queryFlagHelp          = "Fully-qualified system to lookup from all provided DNS servers."
		expectedAnswerFlagHelp = "IP Address expected as the answer from all provided DNS servers."
	)

	config := Config{}

	flag.StringVar(&config.ConfigFile, "config-file", "", "Full path to optional TOML-formatted configuration file. See config.example.toml for a starter template.")
	flag.BoolVar(&config.Version, "version", false, versionFlagHelp)
	flag.BoolVar(&config.Version, "v", false, versionFlagHelp+" (shorthand)")

	flag.StringVar(&config.Query, "query", "", queryFlagHelp)
	flag.StringVar(&config.Query, "q", "", queryFlagHelp)

	flag.StringVar(&config.ExpectedAnswer, "expect", "", expectedAnswerFlagHelp)
	flag.StringVar(&config.ExpectedAnswer, "e", "", expectedAnswerFlagHelp)

	fh, err := os.Open(config.ConfigFile)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer fh.Close()

	data, err := ioutil.ReadAll(fh)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if err := toml.Unmarshal(data, config); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	var msg dns.Msg
	fqdn := dns.Fqdn("stacktitan.com")
	msg.SetQuestion(fqdn, dns.TypeA)
	in, err := dns.Exchange(&msg, "8.8.8.8:53")
	if err != nil {
		panic(err)
	}
	if len(in.Answer) < 1 {
		fmt.Println("No records")
		return
	}
	for _, answer := range in.Answer {
		if a, ok := answer.(*dns.A); ok {
			fmt.Println(a.A)
		}
	}
}
