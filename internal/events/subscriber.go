// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package events

import (
	"context"
	"fmt"

	"github.com/rs/xid"
)

// Subscriber is implemented by pubsub implementations and supports subscribing to events.
type Subscriber interface {
	// On subscribes to an event until the listener is unsubscibed either internally or by calling Subscriber.Unsubscribe()
	// with the Listener or calling Listener.Unsubscribe() directly.
	On(ctx context.Context, event Event, buffer int) *Listener

	// Once works similarly to On, however nlike On, when Once receives its first event the listener will be closed automatically.
	Once(ctx context.Context, event Event) *Listener

	// Unsubscribe unsubscribes the given Listener by calling the Listener.Unsubscribe() method.
	Unsubscribe(lis *Listener)
}

func (p *pubSub) On(ctx context.Context, event Event, buffer int) *Listener {
	return p.on(ctx, event, buffer)
}

func (p *pubSub) Once(ctx context.Context, event Event) *Listener {
	return p.once(ctx, event)
}

func (p *pubSub) Unsubscribe(lis *Listener) {
	p.unsubscrbe(lis)
}

func (p *pubSub) on(ctx context.Context, event Event, buffer int) *Listener {
	id, recv := p.bus.On(string(event), buffer)

	l := &Listener{
		impl:  p,
		ctx:   ctx,
		recv:  recv,
		Event: event,
		ID:    id.String(),
	}

	l.init()

	return l
}

func (p *pubSub) once(ctx context.Context, event Event) *Listener {
	id, recv := p.bus.Once(string(event))

	l := &Listener{
		impl:  p,
		ctx:   ctx,
		recv:  recv,
		Event: event,
		ID:    id.String(),
	}

	l.init()

	return l
}

func (p *pubSub) unsubscrbe(lis *Listener) error {
	id, err := xid.FromString(lis.ID)
	if err != nil {
		return fmt.Errorf("parsing id: %w", err)
	}

	p.bus.Unsubscribe(string(lis.Event), id)

	return nil
}
