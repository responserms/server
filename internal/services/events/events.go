// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package events

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/responserms/server/internal/log"
	"github.com/responserms/server/internal/services/events/bus"
	"gocloud.dev/pubsub"
	"golang.org/x/sync/semaphore"
)

// Event is the type of event being published.
type Event string

// PubSub is a Publisher/Subscriber implementation.
type PubSub interface {
	Publisher
	Subscriber

	// Shutdown should be called once before the program exits to close all existing connections
	// and finish processing any events already in-flight.
	Close(ctx context.Context)
}

// pubSub is a PubSub
type pubSub struct {
	bus *bus.Bus
	pub *pubsub.Topic
	sub *pubsub.Subscription
	sem *semaphore.Weighted
	log log.Logger

	// driver-specific connections
	natsConn *nats.Conn
}

// Options configures all of the options for the PubSub implementation.
type Options struct {
	Logger  log.ComponentLogger
	Driver  Driver
	Weight  int64
	Subject string
	URL     string
}

// New creates a new PubSub implementation that can be used to publish and subscribe to events
// in Response.
func New(ctx context.Context, options *Options) (PubSub, error) {
	var ps = &pubSub{}

	// logger
	ps.log = options.Logger.Component("events")

	// create an event bus with a Component logger for bus, this will make the logs prepend with pubsub.bus
	ps.bus = bus.New(ps.log)

	ps.log.Info("starting", log.Attributes{
		"driver":  string(options.Driver),
		"subject": options.Subject,
		"url":     options.URL,
	})

	// configure according to our driver implementations
	switch options.Driver {
	case NATS:
		if err := ps.setupNATSDriver(options.URL, options.Subject); err != nil {
			return nil, err
		}
	default:
		// Here we emit a warning that the user is using a driver that does not support clustering of events.
		// Basically, what this means it that any events emitted/listened for in this Response application instance
		// will not be shared with others and other instances will not share their events with this instance. This is
		// not typically a problem if they know what they are doing, though it's best to use a real pubsub implementation
		// in production.
		ps.log.Warn("using the embedded driver for single-server clusters only")

		if err := ps.setupEmbeddedDriver(ctx, options.Subject); err != nil {
			return nil, fmt.Errorf("configuring embedded driver: %w", err)
		}
	}

	// if no weight specified, use default of 10
	weight := options.Weight
	if weight == 0 {
		weight = 10
	}

	// initialize the weighted semaphore
	ps.sem = semaphore.NewWeighted(weight)

	// start a goroutine that processes the received messages from the pubsub implementation
	go ps.processMessages(ctx)

	ps.log.Debug("ready")

	return ps, nil
}

func (p *pubSub) Close(ctx context.Context) {
	// shutdown the publisher implementations
	if err := p.pub.Shutdown(ctx); err != nil {
		p.log.Error("publisher failed to properly shutdown", log.Attributes{
			"error": err,
		})
	}

	p.log.Trace("publisher shutdown")

	// shutdown the subscriber implementations
	if err := p.sub.Shutdown(ctx); err != nil {
		p.log.Error("subscruber failed to properly shutdown", log.Attributes{
			"error": err,
		})
	}

	p.log.Trace("subscriber shutdown")

	// close the internal event bus
	p.bus.Close()
	p.log.Info("bus closed")

	// close the nats connection
	if p.natsConn != nil {
		p.natsConn.Close()
		p.log.Trace("nats connection closed")
	}
}
