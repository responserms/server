// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package events

import (
	"context"

	"github.com/responserms/server/internal/log"
	"gocloud.dev/pubsub"
)

// Publisher is implemented by pubsub implementations that can publish events.
type Publisher interface {
	// Publish publishes events to the Response cluster using the configured pubsub implementation.
	Publish(ctx context.Context, messages ...*pubsub.Message) error
}

// Publish publishes events to the Response cluster using the configured pubsub implementation.
func (p *pubSub) Publish(ctx context.Context, messages ...*pubsub.Message) error {
	for _, msg := range messages {
		go func(msg *pubsub.Message) {
			if err := p.pub.Send(context.Background(), msg); err != nil {
				p.log.Error("failed to publish message", log.Attributes{
					"event": msg.Metadata["event"],
				})
			}
		}(msg)
	}

	return nil
}
