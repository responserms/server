// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package log

// Attributes allows configuring the additional attributes logged with the
// log interface.
type Attributes map[string]interface{}

// ToMethodFormat converts the Attributes into the format expected by the
// internal logging implementation.
func (a *Attributes) ToMethodFormat() []interface{} {
	var s []interface{}

	for k, v := range *a {
		s = append(s, k, v)
	}

	return s
}

// mergeManyAttributes merges all of the provided attributes into a single slice
// of interface{} for providing to a logger.
func mergeManyAttributes(a ...Attributes) []interface{} {
	var merged []interface{}

	for _, v := range a {
		merged = append(merged, v.ToMethodFormat()...)
	}

	return merged
}
