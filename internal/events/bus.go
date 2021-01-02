// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package events

import (
	"context"

	"github.com/responserms/server/internal/log"
)

func (p *pubSub) processMessages(ctx context.Context) {
	for {
		msg, err := p.sub.Receive(ctx)
		if err != nil {
			p.log.Error("failed to receive message", log.Attributes{
				"error": err.Error(),
			})

			break
		}

		err = p.sem.Acquire(ctx, 1)
		if err != nil {
			p.log.Error("failed to acquire semaphore", log.Attributes{
				"error": err.Error(),
			})

			break
		}

		go func() {
			defer p.sem.Release(1)
			defer msg.Ack()

			err := p.bus.Publish(msg.Metadata["event"], msg)
			if err != nil {
				p.log.Error("failed to publish event", log.Attributes{
					"event": msg.Metadata["event"],
					"error": err.Error(),
				})
			}
		}()
	}
}
