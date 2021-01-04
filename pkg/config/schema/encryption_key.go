package schema

import (
	"encoding/base64"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/responserms/spec/parser"
	"github.com/zclconf/go-cty/cty"
)

// verify we meet the interface definition
var _ parser.BlockDefinition = (*encryptionKeyDef)(nil)
var _ parser.NamedBlockDefinition = (*encryptionKeyDef)(nil)

// Problem messages, use for comparison if necessary.
const (
	// encoding problems
	DiagEncryptionKeyEncoding       = "The encryption key provided is not base64 encoded"
	DiagEncryptionKeyEncodingDetail = "Encryption keys must be base64 encoded to prevent the need for escaping when passing the key in via" +
		"environment variables and/or config files."

	// length problems
	DiagEncryptionKeyBelowLength  = "The encryption key provided is less than 32 bytes in length"
	DiagEncryptionKeyAboveLength  = "The encryption key provided is more than 32 bytes in length"
	DiagEncryptionKeyLengthDetail = "Encryption keys must be exactly 32 bytes in length to enable the encryption algorithm. You can use" +
		"the 'response keygen' command to generate a suitable encryption key."

	// Bytes required for AES 256-bit encryption algorithm
	EncryptionByteRequirement = 32
)

type encryptionKeyDef struct{}

func (d *encryptionKeyDef) Name() string {
	return "encryption_key"
}

func (d *encryptionKeyDef) Spec() hcldec.Spec {
	return &hcldec.ValidateSpec{
		Wrapped: &hcldec.AttrSpec{
			Name:     "encryption_key",
			Type:     cty.String,
			Required: true,
		},
		Func: func(value cty.Value) hcl.Diagnostics {
			diag := hcl.Diagnostics{}

			if value.IsNull() {
				return hcl.Diagnostics{}
			}

			key, err := base64.StdEncoding.DecodeString(value.AsString())
			if err != nil {
				diag = append(diag, &hcl.Diagnostic{
					Severity: hcl.DiagError,
					Summary:  DiagEncryptionKeyEncoding,
					Detail:   DiagEncryptionKeyEncodingDetail,
				})

				// not valid anyway, fail early
				return diag
			}

			if len(key) < EncryptionByteRequirement {
				diag = append(diag, &hcl.Diagnostic{
					Severity: hcl.DiagError,
					Summary:  DiagEncryptionKeyBelowLength,
					Detail:   DiagEncryptionKeyLengthDetail,
				})
			}

			if len(key) > EncryptionByteRequirement {
				diag = append(diag, &hcl.Diagnostic{
					Severity: hcl.DiagError,
					Summary:  DiagEncryptionKeyAboveLength,
					Detail:   DiagEncryptionKeyLengthDetail,
				})
			}

			return diag
		},
	}
}
