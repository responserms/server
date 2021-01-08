// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package events

import (
	"context"
	"fmt"

	"gocloud.dev/pubsub"

	// used for embedded driver.
	_ "gocloud.dev/pubsub/mempubsub"
)

const (
	// Embedded is an embedded pubsub driver that provides in-memory event subscribing and publishing
	// and is acceptable for testing and development purposes. In general you should use a more robust
	// pubsub implementation in production.
	Embedded Driver = "embedded"
)

// setupEmbeddedDriver configures the embedded driver.
func (p *pubSub) setupEmbeddedDriver(ctx context.Context, subject string) error {
	pub, err := pubsub.OpenTopic(ctx, fmt.Sprintf("mem://%s", subject))
	if err != nil {
		return fmt.Errorf("start publisher: %w", err)
	}

	p.log.Trace("started publisher")

	sub, err := pubsub.OpenSubscription(ctx, fmt.Sprintf("mem://%s", subject))
	if err != nil {
		return fmt.Errorf("start subscriber: %w", err)
	}

	p.log.Trace("started subscriber")

	p.pub = pub
	p.sub = sub

	return nil
}
