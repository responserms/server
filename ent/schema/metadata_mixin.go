// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/mixin"
)

type MetadataMixin struct {
	mixin.Schema
}

// Fields registers one or more fields on the including schema.
func (MetadataMixin) Fields() []ent.Field {
	return nil
}

// Edges registers the metadata edges on the including schema.
func (MetadataMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("metadata", Metadata.Type).
			Unique(),
	}
}
