// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package annotations

import "github.com/facebook/ent/schema"

type Index struct {
	schema.Annotation

	Index string
}

func (i *Index) Name() string {
	return "index"
}
