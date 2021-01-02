// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package schema

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
)

type ipErrorGetter struct{}

func (g *ipErrorGetter) GetPrivateIP() (string, error) {
	return "", errors.New("simulated error")
}

func (g *ipErrorGetter) GetPublicIP() (string, error) {
	return "", errors.New("simulated error")
}

func (g *ipErrorGetter) GetInterfaceIP(namedIf string) (string, error) {
	return "", errors.New("simulated error")
}

func TestNetworkFunctions(t *testing.T) {
	errorGetter := &ipErrorGetter{}

	t.Run("err when returning private ip returns nil", func(t *testing.T) {
		val, err := configureGetPrivateIP(errorGetter).Call([]cty.Value{})

		assert.Contains(t, err.Error(), "simulated error")
		assert.True(t, val.IsNull())
	})

	t.Run("err when returning public ip returns nil", func(t *testing.T) {
		val, err := configureGetPublicIP(errorGetter).Call([]cty.Value{})

		assert.Contains(t, err.Error(), "simulated error")
		assert.True(t, val.IsNull())
	})

	t.Run("err when returning interface ip returns nil", func(t *testing.T) {
		val, err := configureGetInterfaceIP(errorGetter).Call([]cty.Value{cty.StringVal("eth0")})

		assert.Contains(t, err.Error(), "simulated error")
		assert.True(t, val.IsNull())
	})
}
