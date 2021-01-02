// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package config

import (
	"reflect"
	"testing"

	"github.com/matryer/is"
)

func Test_newEmptyConfig(t *testing.T) {
	tests := []struct {
		name string
		want *Config
	}{
		{
			name: "empty config is created",
			want: &Config{
				HTTP: &HTTPConfig{
					TLS: &TLSConfig{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newEmptyConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newEmptyConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_applyDefaults(t *testing.T) {
	type args struct {
		config *Config
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "defaults are applied",
			args: args{
				config: &Config{
					HTTP: &HTTPConfig{},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := applyDefaults(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("applyDefaults() error = %v, wantErr %v", err, tt.wantErr)
			}

			is := is.New(t)
			is.Equal(tt.args.config.HTTP.Port, 8080)
		})
	}
}
