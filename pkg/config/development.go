// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package config

import (
	"github.com/responserms/server/internal/log"
)

const (
	// Encryption key for in-memory mode where data is never persisted.
	inMemoryEncryptionKey = "WQYLN3gBAgYaTe5b7RtZgKw+FpnGAlmlxyQxiLK6YWo="
)

// NewDevelopment creates a Config instance for development purposes where persistence,
// pubsub, and other service requirements are minimal and can rely on the default in-memory
// implementations.
//
// This function returns a working Response Server Config instance but persistence is not
// enabled by default. If persistence is required Response will need to be started with a config
// file instead of with development defaults.
func NewDevelopment() (*Config, error) {
	c := newEmptyConfig()

	// apply defaults
	c = applyDefaults(c)

	// We hardcode an encryption key because development mode will always be in-memory and will not
	// persist any data so it's not necessary to truly care about this.
	c.EncryptionKey = inMemoryEncryptionKey

	// Enable debug mode for logging while running in development mode.
	c.LogLevel = log.Debug

	return c, nil
}
