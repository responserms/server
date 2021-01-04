// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package config

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/responserms/server/internal/log"
)

type AutoTLSDNSConfig struct {
	Provider string `hcl:"solver,label"`
	APIToken string `hcl:"token,attr"`
}

type AutoTLSConfig struct {
	Production bool              `hcl:"production,optional"`
	Email      string            `hcl:"email,attr"`
	Domains    []string          `hcl:"domains,attr"`
	DNS        *AutoTLSDNSConfig `hcl:"dns,block"`
}

type TLSConfig struct {
	Port     int            `hcl:"port,optional"`
	CertPath string         `hcl:"cert_path,optional"`
	KeyPath  string         `hcl:"key_path,optional"`
	Auto     *AutoTLSConfig `hcl:"auto,block"`
}

type HTTPConfig struct {
	BindAddress   string     `hcl:"bind_address,optional"`
	Port          int        `hcl:"port,optional"`
	MaxUploadSize int        `hcl:"max_upload_size,optional"`
	TLS           *TLSConfig `hcl:"tls,block"`
}

type EventsConfig struct {
	Type    string `hcl:"type,optional"`
	URL     string `hcl:"url,optional"`
	Subject string `hcl:"subject,optional"`
}

type DatabaseConfig struct {
	Type     string            `hcl:"type,optional"`
	Path     string            `hcl:"path,optional"`
	URL      string            `hcl:"url,optional"`
	Name     string            `hcl:"name,optional"`
	Host     string            `hcl:"host,optional"`
	Port     int               `hcl:"port,optional"`
	Username string            `hcl:"username,optional"`
	Password string            `hcl:"password,optional"`
	Options  map[string]string `hcl:"options,optional"`
}

type DeveloperConfig struct {
	Profiling bool `hcl:"profiling,optional"`
}

type ClusterAutoJoinConfig struct {
	Type     string            `hcl:"type,label"`
	URL      string            `hcl:"url,optional"`
	Subject  string            `hcl:"subject,optional"`
	Provider string            `hcl:"provider,optional"`
	args     map[string]string `hcl:"args,optional"`
}

type ClusterConfig struct {
	Join              *ClusterAutoJoinConfig `hcl:"autojoin,block"`
	Environment       string                 `hcl:"environment,attr"`
	BindAddress       string                 `hcl:"bind_address,optional"`
	BindPort          int                    `hcl:"bind_port,optional"`
	BindMemberAddress string                 `hcl:"memberlist_bind_address,optional"`
	BindMemberPort    int                    `hcl:"memberlist_bind_port,optional"`
}

type Config struct {
	EncryptionKey string `hcl:"encryption_key,optional"`
	LogLevel      log.Level
	Database      *DatabaseConfig  `hcl:"database,block"`
	HTTP          *HTTPConfig      `hcl:"http,block"`
	Events        *EventsConfig    `hcl:"events,block"`
	Cluster       *ClusterConfig   `hcl:"cluster,block"`
	Developer     *DeveloperConfig `hcl:"developer,block"`
	Remaining     hcl.Body         `hcl:",remain"`
}

// newEmptyConfig creates a new empty configuration with embedded structs initialized. This is a requirement
// for the defaults package to properly configure all of the default values.
func newEmptyConfig() *Config {
	config := &Config{
		Database:  &DatabaseConfig{},
		HTTP:      &HTTPConfig{},
		Cluster:   &ClusterConfig{},
		Events:    &EventsConfig{},
		Developer: &DeveloperConfig{},
	}

	return config
}

// SetLogLevelFromStr sets the log level from the string-variant of the level's name. Supported strings include
// trace, debug, info, warn, error, and off. If an unsupported string is provided the `info` level will be set.
func (c *Config) SetLogLevelFromStr(level string) {
	switch level {
	case "trace":
		c.LogLevel = log.Trace
	case "debug":
		c.LogLevel = log.Debug
	case "info":
		c.LogLevel = log.Info
	case "warn":
		c.LogLevel = log.Warn
	case "error":
		c.LogLevel = log.Error
	case "off":
		c.LogLevel = log.Off
	default:
		c.LogLevel = log.Info
	}
}
