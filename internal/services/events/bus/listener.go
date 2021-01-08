// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package bus

import (
	"errors"
	"sync"

	"github.com/rs/xid"
	"gocloud.dev/pubsub"
)

var (
	// ErrListenerIsClosed is returned when a listener is closed and a Send is attempted
	ErrListenerIsClosed = errors.New("listener is closed")
)

// Listener tracks a single subscription.
type Listener struct {
	id xid.ID
	ch chan *pubsub.Message
	mu sync.RWMutex

	once   bool
	count  int
	closed bool
}

// newListener creates a new Listener, returning a pointer to it. Passing true to once
// will self-destruct the listener once it has received exactly one event. No further
// evnets will be allowed.
func (b *Bus) newListener(once bool, buffer int) *Listener {
	return &Listener{
		id:   xid.New(),
		ch:   make(chan *pubsub.Message, buffer),
		once: once,
	}
}

// Close closes the channel and prevents further message sending.
func (l *Listener) Close() {
	close(l.ch)
	l.closed = true
}

// Send sends the provided msg through the listener's channel, incrementing the message
// count and
func (l *Listener) Send(msg *pubsub.Message) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.closed {
		return ErrListenerIsClosed
	}

	l.ch <- msg
	l.count++

	if l.once && l.count == 1 {
		l.Close()
	}

	return nil
}

// Count returns the total count of messages receives through the Listener.
func (l *Listener) Count() int {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.count
}

// HasReceivedMessages returns true if the Listener has receives at least
// once message.
func (l *Listener) HasReceivedMessages() bool {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.count > 0
}
