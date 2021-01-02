package schema

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/responserms/spec/parser"
	"github.com/zclconf/go-cty/cty"
)

// verify we meet the interface definition
var _ parser.BlockDefinition = (*httpDef)(nil)
var _ parser.NamedBlockDefinition = (*httpDef)(nil)

type httpDef struct{}

func (d *httpDef) Name() string {
	return "http"
}

func (d *httpDef) Spec() hcldec.Spec {
	return &hcldec.ObjectSpec{
		"bind_address": &hcldec.AttrSpec{
			Name: "bind_address",
			Type: cty.String,
		},
		"port": &hcldec.AttrSpec{
			Name: "port",
			Type: cty.Number,
		},
		"max_upload_size": &hcldec.AttrSpec{
			Name: "max_upload_size",
			Type: cty.String,
		},
		"tls": &hcldec.BlockSpec{
			TypeName: "tls",
			Nested: &hcldec.ValidateSpec{
				Wrapped: &hcldec.ObjectSpec{
					"port": &hcldec.AttrSpec{
						Name: "port",
						Type: cty.Number,
					},
					"cert_path": &hcldec.AttrSpec{
						Name: "cert_path",
						Type: cty.String,
					},
					"key_path": &hcldec.AttrSpec{
						Name: "key_path",
						Type: cty.String,
					},
					"auto": &hcldec.BlockSpec{
						TypeName: "auto",
						Nested: &hcldec.ObjectSpec{
							"production": &hcldec.AttrSpec{
								Name: "production",
								Type: cty.Bool,
							},
							"email": &hcldec.AttrSpec{
								Name:     "email",
								Type:     cty.String,
								Required: true,
							},
							"domains": &hcldec.AttrSpec{
								Name:     "domains",
								Type:     cty.Set(cty.String),
								Required: true,
							},
							"dns": &hcldec.BlockSpec{
								TypeName: "dns",
								Nested: &hcldec.ObjectSpec{
									"solver": &hcldec.ValidateSpec{
										Wrapped: &hcldec.BlockLabelSpec{
											Index: 0,
											Name:  "solver",
										},
										Func: func(value cty.Value) hcl.Diagnostics {
											var allowedSolvers = []string{"cloudflare", "vultr", "digitalocean"}

											if value.IsNull() {
												return hcl.Diagnostics{}
											}

											valStr := value.AsString()
											for _, solver := range allowedSolvers {
												if solver == valStr {
													return hcl.Diagnostics{}
												}
											}

											return hcl.Diagnostics{
												{
													Severity: hcl.DiagError,
													Summary:  fmt.Sprintf("The provided solver (%q) is not a valid solver type", valStr),
													Detail:   fmt.Sprintf("The solver type must be one of: %s", strings.Join(allowedSolvers, ", ")),
												},
											}
										},
									},
									"token": &hcldec.AttrSpec{
										Name:     "token",
										Type:     cty.String,
										Required: true,
									},
								},
							},
						},
					},
				},
				Func: func(value cty.Value) hcl.Diagnostics {
					if value.IsNull() {
						return hcl.Diagnostics{}
					}

					v := value.AsValueMap()

					if !v["cert_path"].IsNull() && (v["key_path"].IsNull() || v["key_path"].AsString() == "") {
						return hcl.Diagnostics{
							{
								Severity: hcl.DiagError,
								Summary:  "The `key_path` attribute must be set when `cert_path` is set",
								Detail: "You provided a `cert_path` but did not provide the corresponding `key_path` attribute. Set the" +
									"`key_path` attribute equal to the relative or absolute path to your certificate's key file.",
							},
						}
					}

					if !v["key_path"].IsNull() && (v["cert_path"].IsNull() || v["cert_path"].AsString() == "") {
						return hcl.Diagnostics{
							{
								Severity: hcl.DiagError,
								Summary:  "The `cert_path` attribute must be set when `key_path` is set",
								Detail: "You provided a `key_path` but did not provide the corresponding `cert_path` attribute. Set the" +
									"`cert_path` attribute equal to the relative or absolute path to your certificate file.",
							},
						}
					}

					return hcl.Diagnostics{}
				},
			},
		},
	}
}
