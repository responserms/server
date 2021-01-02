// Copyright (c) 2020 Contaim, LLC
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

type databaseDef struct{}

// Public diagnostic summaries
const (
	DiagUnsupportedDatabaseDriver = "Unsupported database driver"
)

func (s *databaseDef) Name() string {
	return "database"
}

func (s *databaseDef) Spec() hcldec.Spec {
	return &hcldec.BlockSpec{
		TypeName: "database",
		Required: false,
		Nested: &hcldec.ObjectSpec{
			"type": &hcldec.ValidateSpec{
				Wrapped: &hcldec.AttrSpec{
					Name:     "type",
					Type:     cty.String,
					Required: true,
				},
				Func: func(value cty.Value) hcl.Diagnostics {
					validDrivers := []string{"embedded", "sqlite", "postgres", "mysql"}

					if value.IsNull() {
						// fallback to default validation of a block label being required
						return hcl.Diagnostics{}
					}

					driver := value.AsString()
					for _, valid := range validDrivers {
						if valid == driver {
							return hcl.Diagnostics{}
						}
					}

					validDriverStr := strings.Join(validDrivers, ", ")
					return hcl.Diagnostics{
						{
							Summary: DiagUnsupportedDatabaseDriver,
							Detail:  fmt.Sprintf("The %s driver is not a supported driver. Must be one of: %s", driver, validDriverStr),
						},
					}
				},
			},
			"url": &hcldec.AttrSpec{
				Name:     "url",
				Type:     cty.String,
				Required: false,
			},
			"host": &hcldec.AttrSpec{
				Name:     "host",
				Type:     cty.String,
				Required: false,
			},
			"port": &hcldec.AttrSpec{
				Name:     "port",
				Type:     cty.String,
				Required: false,
			},
			"database": &hcldec.AttrSpec{
				Name:     "database",
				Type:     cty.String,
				Required: false,
			},
			"username": &hcldec.AttrSpec{
				Name:     "username",
				Type:     cty.String,
				Required: false,
			},
			"password": &hcldec.AttrSpec{
				Name:     "password",
				Type:     cty.String,
				Required: false,
			},
			"options": &hcldec.BlockAttrsSpec{
				TypeName:    "options",
				ElementType: cty.String,
				Required:    false,
			},
		},
	}
}
