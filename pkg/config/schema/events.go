// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package schema

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

type eventsDef struct{}

// Diag messages
var (
	// Diagnostic messages regarding an invalid message bus implementation type
	DiagInvalidPubSubType       = "The events type is not valid"
	DiagInvalidPubSubTypeDetail = "The events type must be one of: embedded, nats"

	DiagNATSRequiresURL       = "The 'nats' events type requires a URL."
	DiagNatsRequiresURLDetail = "You must provide the 'url' attribute when using the 'nats' pubsub type."
)

func (s *eventsDef) Name() string {
	return "events"
}

func (s *eventsDef) Spec() hcldec.Spec {
	return &hcldec.BlockSpec{
		TypeName: s.Name(),
		Required: false,
		Nested: &hcldec.ValidateSpec{
			Wrapped: &hcldec.ObjectSpec{
				"type": &hcldec.ValidateSpec{
					Wrapped: &hcldec.AttrSpec{
						Name:     "type",
						Type:     cty.String,
						Required: true,
					},
					Func: func(value cty.Value) hcl.Diagnostics {
						if value.IsNull() {
							// fallback to default "required" setting
							return hcl.Diagnostics{}
						}

						validPubSub := []string{"nats", "embedded"}
						selectedType := value.AsString()
						for _, d := range validPubSub {
							if selectedType == d {
								return hcl.Diagnostics{}
							}
						}

						return hcl.Diagnostics{
							{
								Severity: hcl.DiagError,
								Summary:  DiagInvalidPubSubType,
								Detail:   DiagInvalidPubSubTypeDetail,
							},
						}
					},
				},
				"subject": &hcldec.AttrSpec{
					Name:     "subject",
					Type:     cty.String,
					Required: false,
				},
				"url": &hcldec.AttrSpec{
					Name:     "url",
					Type:     cty.String,
					Required: false,
				},
			},
			Func: func(value cty.Value) hcl.Diagnostics {
				return hcl.Diagnostics{}
			},
		},
	}
}
