// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package schema

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty/function"
)

// NewContext creates a new hcl.EvalContext and returns the pointer to be used
// with the schema. This allows the caller to automatically insert expected functions
// and variables.
func NewContext() *hcl.EvalContext {
	ipGetter := &IPGetter{}

	return &hcl.EvalContext{
		Functions: map[string]function.Function{
			"env":          configureEnvOrDefault(),
			"private_ip":   configureGetPrivateIP(ipGetter),
			"public_ip":    configureGetPublicIP(ipGetter),
			"interface_ip": configureGetInterfaceIP(ipGetter),
		},
	}
}
