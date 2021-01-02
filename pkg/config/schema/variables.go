// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package schema

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/responserms/spec/parser"
	"github.com/zclconf/go-cty/cty"
)

// validate that we meet the interface requirements
var _ parser.BlockDefinition = (*variablesDef)(nil)
var _ parser.VariableInjector = (*variablesDef)(nil)
var _ parser.NamedBlockDefinition = (*variablesDef)(nil)

type variablesDef struct{}

func (d *variablesDef) Name() string {
	return "variables"
}

func (d *variablesDef) Spec() hcldec.Spec {
	return &hcldec.BlockAttrsSpec{
		TypeName:    d.Name(),
		ElementType: cty.String,
		Required:    false,
	}
}

func (d *variablesDef) Variables(v cty.Value) parser.InjectableVariables {
	if v.IsNull() {
		return nil
	}

	return v.AsValueMap()
}
