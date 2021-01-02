package schema_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/responserms/server/pkg/config/schema"
	"github.com/responserms/spec"
	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
)

func TestVars(tt *testing.T) {
	tt.Run("ensure vars block is parsed first", func(t *testing.T) {
		s := spec.New(schema.Schema())

		// vars needs to be processed first so that we can use them in other blocks,
		// this test simply ensures that the `vars` blocks are processed first.
		assert.Equal(t, "variables", s.Registrations()[0].BlockName)
	})

	tests := []struct {
		file         string
		expectPass   bool
		expectErrMsg string
	}{
		// vars block can be used only once
		{
			file:         "./testdata/vars/duplicate.hcl",
			expectErrMsg: "Duplicate variables block",
		},

		// vars is not required
		{
			file:       "./testdata/empty.hcl",
			expectPass: true,
		},
	}

	for i, test := range tests {
		tt.Run(fmt.Sprintf("%02d", i), func(t *testing.T) {
			s := spec.NewSubset(schema.VariablesDefinition())

			// testadata/vars.hcl does not contain the encryption_key attribute
			file := s.ParseHCLFile(test.file)
			assert.False(t, file.HasErrors())

			diags := s.Parse(&hcl.EvalContext{})

			if test.expectPass {
				assert.False(t, diags.HasErrors())

				return
			}

			assert.True(t, diags.HasErrors())
			assert.Contains(t, diags.Error(), test.expectErrMsg)
		})
	}

	tt.Run("ensure all provided vars become variables in the context", func(t *testing.T) {
		s := spec.NewSubset(schema.VariablesDefinition())
		f := s.Files("./testdata/vars/valid.hcl")

		// we expect no problems when parsing our test file
		assert.False(t, f.HasErrors())

		ctx := &hcl.EvalContext{}
		s.Parse(ctx)

		isEqual := reflect.DeepEqual(ctx.Variables, map[string]cty.Value{
			"my_var": cty.StringVal("says hello"),
		})

		assert.True(t, isEqual)
	})
}
