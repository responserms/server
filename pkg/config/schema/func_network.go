// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package schema

import (
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

func configureGetPrivateIP(get PrivateIPGetter) function.Function {
	return function.New(&function.Spec{
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			ip, err := get.GetPrivateIP()
			if err != nil {
				return cty.StringVal(""), err
			}

			return cty.StringVal(ip), nil
		},
	})
}

func configureGetPublicIP(get PublicIPGetter) function.Function {
	return function.New(&function.Spec{
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			ip, err := get.GetPublicIP()
			if err != nil {
				return cty.StringVal(""), err
			}

			return cty.StringVal(ip), nil
		},
	})
}

func configureGetInterfaceIP(get InterfaceIPGetter) function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name:         "namedIfRE",
				Type:         cty.String,
				AllowNull:    false,
				AllowUnknown: false,
			},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			ip, err := get.GetInterfaceIP(args[0].AsString())
			if err != nil {
				return cty.StringVal(""), err
			}

			return cty.StringVal(ip), nil
		},
	})
}
