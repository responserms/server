// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package bus

import (
	"errors"
	"sync"

	"github.com/responserms/server/internal/log"
	"github.com/rs/xid"
	"gocloud.dev/pubsub"
)

var (
	// ErrBusIsClosed is returned when the Bus has been closed and a Publish is attempted.
	ErrBusIsClosed = errors.New("bus has closed")
)

// Bus is an event bus responsible for sending/receiving events.
type Bus struct {
	mu        sync.RWMutex
	logger    log.Logger
	listeners map[string][]*Listener
	closed    bool
}

// New creates a Bus returning its pointer.
func New(logger log.Logger) *Bus {
	bus := &Bus{}
	bus.logger = logger.Component("bus")
	bus.listeners = map[string][]*Listener{}

	return bus
}

// On subscribes to a specific topic with the given buffer size.
func (b *Bus) On(topic string, buffer int) (xid.ID, <-chan *pubsub.Message) {
	b.mu.Lock()
	defer b.mu.Unlock()

	lis := b.newListener(false, buffer)
	b.listeners[topic] = append(b.listeners[topic], lis)

	b.logger.Debug("listener joined", log.Attributes{
		"id":     lis.id.String(),
		"event":  topic,
		"buffer": buffer,
		"once":   false,
		"total":  b.Listeners(topic),
	})

	return lis.id, lis.ch
}

// Once subscribes to a specific topic. The listener will be closed immediately after
// at least one message is received.
func (b *Bus) Once(topic string) (xid.ID, <-chan *pubsub.Message) {
	b.mu.Lock()
	defer b.mu.Unlock()

	lis := b.newListener(true, 1)
	b.listeners[topic] = append(b.listeners[topic], lis)

	b.logger.Debug("listener joined", log.Attributes{
		"id":     lis.id.String(),
		"event":  topic,
		"buffer": 1,
		"once":   true,
		"total":  b.Listeners(topic),
	})

	return lis.id, lis.ch
}

// Unsubscribe unsubscribes the listener by ID from the given topic.
func (b *Bus) Unsubscribe(topic string, id xid.ID) {
	b.mu.Lock()
	defer b.mu.Unlock()

	remove := []int{}
	for i, lis := range b.listeners[topic] {
		if lis.id == id {
			lis.Close()
			remove = append(remove, i)
		}
	}

	for l := range remove {
		b.removeListenerAtPos(topic, l)
	}

	b.logger.Debug("listener left", log.Attributes{
		"id":        id.String(),
		"event":     topic,
		"remaining": b.Listeners(topic),
	})
}

func (b *Bus) removeListenerAtPos(topic string, pos int) {
	if ok := b.listeners[topic][pos]; ok != nil {
		b.listeners[topic] = append(b.listeners[topic][:pos], b.listeners[topic][pos+1:]...)
	}
}

// Publish publishes a msg to the given topic.
func (b *Bus) Publish(topic string, msg *pubsub.Message) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if b.closed {
		b.logger.Error("attempted to publish on closed event bus")

		return ErrBusIsClosed
	}

	b.logger.Info("event received", log.Attributes{
		"event":     topic,
		"listeners": b.Listeners(topic),
	})

	// emit the msg on all listeners
	for i, lis := range b.listeners[topic] {
		if b.logger.IsTrace() {
			b.logger.Trace("publish to listener", log.Attributes{
				"event":     topic,
				"id":        lis.id.String(),
				"remaining": (b.Listeners(topic) - 1) - i,
			})
		}

		err := lis.Send(msg)

		if err == ErrListenerIsClosed {
			b.logger.Error("pushing to listener failed", log.Attributes{
				"id":     lis.id.String(),
				"reason": "closed",
			})

			b.removeListenerAtPos(topic, i)
		}
	}

	// remove any closed listeners
	for i, lis := range b.listeners[topic] {
		if lis.closed {
			b.removeListenerAtPos(topic, i)

			b.logger.Info("listener evicted", log.Attributes{
				"event":  topic,
				"id":     lis.id.String(),
				"reason": "closed",
			})
		}
	}

	return nil
}

// Close closes all listener channels and empties the list of subscribed listeners.
func (b *Bus) Close() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.logger.Info("closing event bus")

	if !b.closed {
		b.closed = true

		for topic, subs := range b.listeners {
			for _, lis := range subs {
				lis.Close()
			}

			b.listeners[topic] = []*Listener{}
		}
	}
}

// Listeners returns the number of listeners that are registered for a given topic.
func (b *Bus) Listeners(topic string) int {
	return len(b.listeners[topic])
}
