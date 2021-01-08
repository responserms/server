// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package bus

import (
	"sync"
	"testing"

	"github.com/responserms/server/internal/log"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"gocloud.dev/pubsub"
)

var logger, _ = log.New(log.WithLogLevel(log.Trace))

func TestNew(t *testing.T) {
	b := New(logger)
	assert.IsType(t, (*Bus)(nil), b)
}

func TestListeners(tt *testing.T) {
	b := New(logger)
	defer b.Close()

	tt.Run("no listeners exist when instantiated", func(t *testing.T) {
		assert.Equal(t, 0, len(b.listeners))
	})

	tt.Run("listener is not closed when created", func(t *testing.T) {
		lis := b.newListener(true, 1)

		assert.Equal(t, false, lis.closed)
	})

	tt.Run("subscribing using On will receive an event on publish", func(t *testing.T) {
		lisID, ch := b.On("tests", 1)
		defer b.Unsubscribe("tests", lisID)

		assert.IsType(t, (xid.ID)(xid.New()), lisID)
		assert.IsType(t, (<-chan *pubsub.Message)(nil), ch)

		want := []byte("ok")
		got := []byte{}

		w := sync.WaitGroup{}
		w.Add(1)

		// listen for the event, push it to got
		go func() {
			defer w.Done()

			ev := <-ch
			got = ev.Body
		}()

		b.Publish("tests", &pubsub.Message{
			Body: want,
		})

		w.Wait()
		assert.Equal(t, want, got)
	})

	tt.Run("subscribing with Once will receive an event and no longer be subscribed", func(t *testing.T) {
		lisID, ch := b.Once("tests.once")

		assert.IsType(t, (xid.ID)(xid.New()), lisID)
		assert.IsType(t, (<-chan *pubsub.Message)(nil), ch)

		want := []byte("ok")
		got := []byte{}

		var w sync.WaitGroup
		w.Add(1)

		// listen for the event, push it to got
		go func() {
			defer w.Done()

			ev := <-ch
			got = ev.Body
		}()

		b.Publish("tests.once", &pubsub.Message{
			Body: want,
		})

		w.Wait()

		assert.Equal(t, want, got)
		assert.Equal(t, 0, b.Listeners("tests.once"))
	})

	tt.Run("subscribing with On and then Unsubscribing does not receive an event", func(t *testing.T) {
		lisID, ch := b.On("tests.unsubscribe", 1)
		b.Unsubscribe("tests.unsubscribe", lisID)

		// ok is false if the channel is closed, which it should be due to the Unsubscribe
		_, ok := (<-ch)
		assert.False(t, ok)
	})

	tt.Run("closed listener is removed upon error", func(t *testing.T) {
		closedListener := b.newListener(false, 1)
		closedListener.Close()

		b.listeners = map[string][]*Listener{
			"test": make([]*Listener, 0),
		}

		b.listeners["test"] = append(b.listeners["test"], closedListener)

		assert.Equal(t, 1, b.Listeners("test"))

		// publish, even though our listener is closed
		b.Publish("test", &pubsub.Message{})

		// the listener should have been removed because it was closed
		assert.Equal(t, 0, b.Listeners("test"))
	})
}

func TestBus_Close(tt *testing.T) {
	tt.Run("Publish() on closed bus errors", func(t *testing.T) {
		b := New(logger)
		b.Close()

		err := b.Publish("test", &pubsub.Message{})
		assert.Equal(t, ErrBusIsClosed, err)
	})

	tt.Run("Close() closes registered listeners", func(t *testing.T) {
		b := New(logger)
		l := b.newListener(false, 1)

		b.listeners["test"] = append(b.listeners["test"], l)
		assert.Equal(t, 1, b.Listeners("test"))
		assert.False(t, l.closed)

		b.Close()
		_, ok := (<-l.ch)

		// channel should be closed, thus ok should be false
		assert.False(t, ok)

		assert.Equal(t, 0, b.Listeners("test"))
		assert.True(t, l.closed)
	})
}
