// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package schema_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/responserms/server/pkg/config/schema"
	"github.com/responserms/spec"
	"github.com/stretchr/testify/assert"
)

func TestEncryptionKey(tt *testing.T) {
	os.Setenv("test_encryption_key", "WQYLN3gBAgYaTe5b7RtZgKw+FpnGAlmlxyQxiLK6YWo=")

	tests := []struct {
		file         string
		expectPass   bool
		expectErrMsg string
	}{
		{
			file:         "./testdata/encryption_key/not_base64.hcl",
			expectErrMsg: schema.DiagEncryptionKeyEncoding,
		},
		{
			file:         "./testdata/encryption_key/too_long.hcl",
			expectErrMsg: schema.DiagEncryptionKeyAboveLength,
		},
		{
			file:         "./testdata/encryption_key/too_short.hcl",
			expectErrMsg: schema.DiagEncryptionKeyBelowLength,
		},
		{
			file:         "./testdata/encryption_key/env_var_no_fallback.hcl",
			expectErrMsg: schema.DiagEncryptionKeyBelowLength,
		},
		{
			file:       "./testdata/encryption_key/valid.hcl",
			expectPass: true,
		},
		{
			file:       "./testdata/encryption_key/env_var.hcl",
			expectPass: true,
		},
	}

	for i, test := range tests {
		tt.Run(fmt.Sprintf("%02d", i), func(t *testing.T) {
			s := spec.NewSubset(schema.EncryptionKeyDefinition())

			file := s.ParseHCLFile(test.file)
			assert.False(t, file.HasErrors())

			diags := s.Parse(schema.NewContext())

			if test.expectPass {
				assert.Falsef(t, diags.HasErrors(), diags.Error())

				return
			}

			assert.True(t, diags.HasErrors())
			assert.Contains(t, diags.Error(), test.expectErrMsg)
		})
	}
}
