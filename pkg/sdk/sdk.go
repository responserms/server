// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package sdk

import (
	"fmt"

	"github.com/responserms/server/ent"
	"github.com/responserms/server/pkg/sdk/event"
)

func Test() {
	user := &ent.User{
		ID: 1,
	}

	fmt.Println(event.UserActivated(user))
}
