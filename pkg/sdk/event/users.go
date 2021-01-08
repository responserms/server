// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package event

import (
	"fmt"

	"github.com/responserms/server/ent"
	"github.com/responserms/server/internal/services/events"
)

// AnyUserRegistered should be emitted whenever a User first registers for Response. The message
// will contain the ent.User that registered as the body.
func AnyUserRegistered() events.Event {
	return events.Event("user.registered")
}

// AnyUserActivated should be emitted whenever any User is activated. The messages will contain
// the ent.User that was activated as the body.
func AnyUserActivated() events.Event {
	return events.Event("user.activated")
}

// UserActivated should be emitted whenever the User is activated. This should only ever be
// emitted once and will never be emitted again unless the user is deactivated and then
// reactivated. The messages will contain the ent.User that was activated as the body.
func UserActivated(user *ent.User) events.Event {
	return events.Event(fmt.Sprintf("user.activated.%d", user.ID))
}

// AnyUserDisabled should be emitted whenever any User is disabled. The message will contain the
// ent.User that was disabled as the body.
func AnyUserDisabled() events.Event {
	return events.Event("user.disabled")
}

// UserDisabled should be emitted whenever the User is disabled. The message will contain the
// ent.User that was disabled as the body.
func UserDisabled(user *ent.User) events.Event {
	return events.Event(fmt.Sprintf("user.disabled.%d", user.ID))
}
