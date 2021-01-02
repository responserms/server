// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package log

import (
	"reflect"
	"testing"
)

func TestAttributes_ToMethodFormat(t *testing.T) {
	tests := []struct {
		name string
		a    *Attributes
		want []interface{}
	}{
		{
			name: "converts to proper format",
			a: &Attributes{
				"key":  "value",
				"key2": "value2",
			},
			want: []interface{}{"key", "value", "key2", "value2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.ToMethodFormat(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attributes.ToMethodFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeManyAttributes(t *testing.T) {
	type args struct {
		a []Attributes
	}

	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "merges many Attributes types",
			args: args{
				a: []Attributes{
					{
						"key": "value",
					},
					{
						"key2": "value2",
					},
				},
			},
			want: []interface{}{"key", "value", "key2", "value2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeManyAttributes(tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeManyAttributes() = %v, want %v", got, tt.want)
			}
		})
	}
}
