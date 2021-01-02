// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package events

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/responserms/server/internal/log"
	"gocloud.dev/pubsub/natspubsub"
)

const (

	// NATS connects to an external NATS cluster to provide event subscribing and publishing across
	// multiple Response instances each connected to the same NATS cluster. This is the preferred
	// production system.
	NATS = "nats"
)

// setupNATSDriver configures the NATS driver.
func (p *pubSub) setupNATSDriver(url, subject string) error {
	natsConn, err := nats.Connect(url, nats.ErrorHandler(func(c *nats.Conn, s *nats.Subscription, e error) {}))
	if err != nil {
		return fmt.Errorf("connect to nats: %w", err)
	}

	p.natsConn = natsConn

	p.log.Info("nats connection established", log.Attributes{
		"server": p.natsConn.ConnectedServerId(),
		"addr":   p.natsConn.ConnectedAddr(),
	})

	pub, err := natspubsub.OpenTopic(p.natsConn, subject, &natspubsub.TopicOptions{})
	if err != nil {
		return fmt.Errorf("start publisher: %w", err)
	}

	p.log.Trace("started publisher")

	sub, err := natspubsub.OpenSubscription(p.natsConn, subject, &natspubsub.SubscriptionOptions{})
	if err != nil {
		return fmt.Errorf("start subscriber: %w", err)
	}

	p.log.Trace("started subscriber")

	p.pub = pub
	p.sub = sub

	return nil
}
