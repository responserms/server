// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package events

import (
	"context"
	"sync"

	"gocloud.dev/pubsub"
)

// Listener is a single subscribed listener for an event.
type Listener struct {
	impl  *pubSub
	once  sync.Once
	ctx   context.Context
	close <-chan struct{}
	recv  <-chan *pubsub.Message

	Event Event
	ID    string
}

func (l *Listener) init() {
	l.once.Do(func() {
		l.watchContext()
	})
}

func (l *Listener) watchContext() {
	go func() {
		<-l.ctx.Done()
		l.Unsubscribe()
	}()
}

// Unsubscribe unsubscribes the listener. The listener will no longer receive
// messages from the subscribed event.
func (l *Listener) Unsubscribe() {
	l.impl.Unsubscribe(l)
}

// Channel returns the channel that emits pubsub.Message's whenever they are
// received. Channel will also
func (l *Listener) Channel() <-chan *pubsub.Message {
	return l.recv
}
