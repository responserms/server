// Copyright (c) 2021 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cluster

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/buraksezer/olric"
	"github.com/responserms/server/internal/log"
)

// Cluster is implemented by services offering clustering.
type Cluster interface {
	// NewStore creates a new distributed key/value store.
	NewStore(name string) (Storer, error)

	// NewLockStore creates a new distributed lock store.
	NewLockStore(name string) (Locker, error)

	// Start starts the cluster service.
	Start() error

	// Shutdown shuts down the cluster service.
	Shutdown(ctx context.Context) error
}

var (
	logExp        = regexp.MustCompile(`^(2.*\s)(\[.*\])`)
	logEmptyBytes = []byte{}
)

type logWriter struct{}

func (l logWriter) Write(bytes []byte) (int, error) {
	log := logExp.FindAllSubmatch(bytes, -1)
	msg := logExp.ReplaceAll(bytes, logExp.Find(logEmptyBytes))

	for i := range log {
		return fmt.Printf(
			"%s %s  server.cluster:%s",
			time.Now().Format("2006-01-02T15:04:05.999-0700"),
			log[i][2],
			msg,
		)
	}

	return 0, nil
}

type cluster struct {
	impl *olric.Olric
}

// New creates a new service implementing the Cluster interface.
func New(log log.ComponentLogger, options *Options) (Cluster, error) {
	err := options.init()
	if err != nil {
		return nil, fmt.Errorf("new cluster config: %w", err)
	}

	logger := log.Component("cluster").StandardLogger()
	logger.SetFlags(0)
	logger.SetOutput(logWriter{})

	options.olric.LogOutput = logWriter{}

	impl, err := olric.New(options.olric)
	if err != nil {
		return nil, fmt.Errorf("new: %s", err)
	}

	return &cluster{
		impl: impl,
	}, nil
}

func (c *cluster) Start() error {
	return c.impl.Start()
}

func (c *cluster) Shutdown(ctx context.Context) error {
	return c.impl.Shutdown(ctx)
}

// func (c *cluster) Members() {
// 	stats, err := c.impl.Stats()
// 	if err != nil {
// 		return nil, fmt.Errorf("stats: %w", err)
// 	}

// 	c.impl.

// 	stats.
// }
