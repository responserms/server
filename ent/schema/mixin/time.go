// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package mixin

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
	"github.com/facebookincubator/ent-contrib/entgql"
)

// Time implements the ent.Mixin for sharing the created_at and updated_at
// fields across multiple schemas.
type Time struct {
	mixin.Schema
}

// Fields returns all fields provided by the mixin.
func (Time) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Annotations(entgql.OrderField("CREATED_AT")).
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Annotations(entgql.OrderField("UPDATED_AT")).
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}
