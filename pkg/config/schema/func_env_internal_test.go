package schema

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
)

func TestTrimQuotes(tt *testing.T) {
	tt.Run("ignores string without quotes", func(t *testing.T) {
		assert.Equal(t, trimQuotes("no quotes"), "no quotes")
	})

	tt.Run("properly removes double quotes", func(t *testing.T) {
		assert.Equal(t, trimQuotes(`"this is quoted"`), "this is quoted")
	})
}

func TestHandleEnvOrDefaultFunc(tt *testing.T) {
	tt.Run("returns proper environment variable", func(t *testing.T) {
		os.Setenv("test_handle_env", "test")
		got, err := configureEnvOrDefault().Call([]cty.Value{cty.StringVal("test_handle_env")})
		assert.Equal(t, nil, err)
		assert.Equal(t, "test", got.AsString())
	})

	tt.Run("returns fallback if no environment variable", func(t *testing.T) {
		got, err := configureEnvOrDefault().Call([]cty.Value{cty.StringVal("test_does_not_exist"), cty.StringVal("fallback")})
		assert.Equal(t, nil, err)
		assert.Equal(t, "fallback", got.AsString())
	})
}
