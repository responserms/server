// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package events

import (
	"encoding/json"

	"gocloud.dev/pubsub"
)

// NewMessage creates a new Message containing the event and body. This message can be
// sent using the Publish method.
func NewMessage(body interface{}, events ...Event) ([]*pubsub.Message, error) {
	var send []byte

	// if already a byte slice no marshaling is required
	if b, ok := body.([]byte); ok {
		send = b
	}

	// if send hasn't already been set
	if len(send) == 0 {
		s, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		send = s
	}

	// send all messages
	messages := []*pubsub.Message{}
	for _, e := range events {
		messages = append(messages, &pubsub.Message{
			Body: send,
			Metadata: map[string]string{
				"event": string(e),
			},
		})
	}

	return messages, nil
}

// Unmarshal unmarshals the *pubsub.Message into the provided interface.
func Unmarshal(msg *pubsub.Message, into interface{}) error {
	return json.Unmarshal(msg.Body, into)
}
