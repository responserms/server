// Copyright (c) 2021 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package schema

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

type clusterDef struct{}

// Public diagnostic summaries
const (
	DiagUnsupportedAutoJoinType = "Unsupported cluster auto-join type"
)

func (s *clusterDef) Name() string {
	return "cluser"
}

func (s *clusterDef) Spec() hcldec.Spec {
	return &hcldec.BlockSpec{
		TypeName: s.Name(),
		Required: false,
		Nested: &hcldec.ObjectSpec{
			"environment": &hcldec.ValidateSpec{
				Wrapped: &hcldec.AttrSpec{
					Name:     "environment",
					Type:     cty.String,
					Required: true,
				},
				Func: func(value cty.Value) hcl.Diagnostics {
					validTypes := []string{"local", "lan", "wan"}

					if value.IsNull() {
						return hcl.Diagnostics{}
					}

					driver := value.AsString()
					for _, valid := range validTypes {
						if valid == driver {
							return hcl.Diagnostics{}
						}
					}

					validTypeStr := strings.Join(validTypes, ", ")
					return hcl.Diagnostics{
						{
							Summary: DiagUnsupportedAutoJoinType,
							Detail:  fmt.Sprintf("The %s environment is not a supported cluster environment type. Must be one of: %s", driver, validTypeStr),
						},
					}
				},
			},
			"bind_address": &hcldec.AttrSpec{
				Name: "bind_address",
				Type: cty.String,
			},
			"bind_port": &hcldec.AttrSpec{
				Name: "bind_port",
				Type: cty.Number,
			},
			"discovery_bind_address": &hcldec.AttrSpec{
				Name: "discovery_bind_address",
				Type: cty.String,
			},
			"discovery_bind_port": &hcldec.AttrSpec{
				Name: "discovery_bind_port",
				Type: cty.Number,
			},
			"members": &hcldec.AttrSpec{
				Name: "members",
				Type: cty.Set(cty.String),
			},
			"autojoin": &hcldec.BlockSpec{
				TypeName: "autojoin",
				Nested: &hcldec.ObjectSpec{
					"type": &hcldec.ValidateSpec{
						Wrapped: &hcldec.BlockLabelSpec{
							Name:  "type",
							Index: 0,
						},
						Func: func(value cty.Value) hcl.Diagnostics {
							validTypes := []string{"nats"}

							if value.IsNull() {
								// fallback to default validation of a block label being required
								return hcl.Diagnostics{}
							}

							driver := value.AsString()
							for _, valid := range validTypes {
								if valid == driver {
									return hcl.Diagnostics{}
								}
							}

							validTypeStr := strings.Join(validTypes, ", ")
							return hcl.Diagnostics{
								{
									Summary: DiagUnsupportedAutoJoinType,
									Detail:  fmt.Sprintf("The %s type is not a supported cluster auto join type. Must be one of: %s", driver, validTypeStr),
								},
							}
						},
					},
					"subject": &hcldec.AttrSpec{
						Name: "subject",
						Type: cty.String,
					},
					"url": &hcldec.AttrSpec{
						Name: "url",
						Type: cty.String,
					},
					"provider": &hcldec.AttrSpec{
						Name: "provider",
						Type: cty.String,
					},
					"args": &hcldec.AttrSpec{
						Name: "args",
						Type: cty.Map(cty.String),
					},
				},
			},
		},
	}
}
