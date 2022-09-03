// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/dnsc
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/apex/log"
	"github.com/pelletier/go-toml/v2"
)

// loadConfigFile is a helper function to handle opening a specified config
// file and importing the settings for use
func (c *Config) loadConfigFile(configFile string) (bool, error) {
	// load config file
	log.WithFields(log.Fields{
		"config_file": configFile,
	}).Debug("Attempting to open config file")

	fh, err := os.Open(filepath.Clean(configFile))
	if err != nil {
		return false, err
	}
	log.Debug("Config file opened")

	// #nosec G307
	// Believed to be a false-positive from recent gosec release
	// https://github.com/securego/gosec/issues/714
	defer func() {
		if err := fh.Close(); err != nil {
			// Ignore "file already closed" errors
			if !errors.Is(err, os.ErrClosed) {
				log.Errorf(
					"loadConfigFile: failed to close file %q: %s",
					configFile,
					err.Error(),
				)
			}
		}
	}()

	log.Debug("Attempting to import config file")
	result, err := c.ImportConfigFile(fh)
	if err != nil {
		return false, err
	}

	return result, err
}

// localConfigFile returns the potential path to a config file found locally
// alongside the executable, or an error if one is encountered constructing
// the path
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

// userConfigFile returns the potential path to a config file found in the
// user's configuration path or an error if one is encountered constructing the
// path
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
	configFile, err := io.ReadAll(fh)
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
