// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package schema

import (
	"os"

	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

func configureEnvOrDefault() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name:         "name",
				Type:         cty.String,
				AllowNull:    false,
				AllowUnknown: false,
			},
		},
		VarParam: &function.Parameter{
			Name:      "default",
			Type:      cty.String,
			AllowNull: true,
		},
		Type: function.StaticReturnType(cty.String),
		Impl: handleEnvOrDefaultFunc,
	})
}

func trimQuotes(s string) string {
	if len(s) >= 2 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
	}
	return s
}

func handleEnvOrDefaultFunc(args []cty.Value, retType cty.Type) (cty.Value, error) {
	envValue := trimQuotes(os.Getenv(args[0].AsString()))

	if len(args) == 1 {
		return cty.StringVal(envValue), nil
	}

	if envValue == "" {
		return args[1], nil
	}

	return cty.StringVal(envValue), nil
}
