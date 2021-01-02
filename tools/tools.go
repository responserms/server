// +build tools

// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package tools

//go:generate go install golang.org/x/tools/cmd/goimports
//go:generate go install github.com/mitchellh/gox
//go:generate go install github.com/hashicorp/go-bindata
//go:generate go install github.com/elazarl/go-bindata-assetfs
//go:generate go install github.com/matryer/drop
//go:generate go install github.com/fatih/gomodifytags
//go:generate go install github.com/99designs/gqlgen
//go:generate go install github.com/facebook/ent/cmd/ent

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/elazarl/go-bindata-assetfs"
	_ "github.com/facebook/ent/cmd/ent"
	_ "github.com/fatih/gomodifytags"
	_ "github.com/hashicorp/go-bindata"
	_ "github.com/matryer/drop"
	_ "github.com/mitchellh/gox"
	_ "golang.org/x/tools/cmd/goimports"
)
