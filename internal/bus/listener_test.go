// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package bus

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"gocloud.dev/pubsub"
)

func TestListener(tt *testing.T) {
	tt.Run("closed listener does not allow sending", func(t *testing.T) {
		l := New(logger).newListener(false, 1)
		l.Close()

		err := l.Send(&pubsub.Message{})
		assert.Equal(t, ErrListenerIsClosed, err)
	})

	tt.Run("listener fails on second send when once is true", func(t *testing.T) {
		l := New(logger).newListener(true, 1)

		var w sync.WaitGroup

		w.Add(1)
		go func() {
			defer w.Done()
			<-l.ch
		}()

		err := l.Send(&pubsub.Message{})
		assert.Nil(t, err)

		w.Wait()

		err = l.Send(&pubsub.Message{})
		assert.Equal(t, ErrListenerIsClosed, err)
	})

	tt.Run("Count() returns proper message count", func(t *testing.T) {
		l := New(logger).newListener(false, 1)
		defer l.Close()

		var w sync.WaitGroup

		handle := func(lis *Listener) {
			defer w.Done()
			<-lis.ch
		}

		w.Add(1)
		go handle(l)

		l.Send(&pubsub.Message{})
		w.Wait()

		assert.Equal(t, 1, l.Count())

		w.Add(1)
		go handle(l)

		l.Send(&pubsub.Message{})
		w.Wait()

		assert.Equal(t, 2, l.Count())

		w.Add(1)
		go handle(l)

		l.Send(&pubsub.Message{})
		w.Wait()

		assert.Equal(t, 3, l.Count())
	})

	tt.Run("HasReceivedMessages() returns true when a message has been received", func(t *testing.T) {
		l := New(logger).newListener(true, 1)
		assert.False(t, l.HasReceivedMessages())

		l.Send(&pubsub.Message{})

		var r bool
		var w sync.WaitGroup

		w.Add(1)
		go func() {
			defer w.Done()

			<-l.ch
			r = true
		}()

		w.Wait()

		assert.True(t, r)
		assert.True(t, l.HasReceivedMessages())
	})
}
