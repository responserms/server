// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package schema

import "github.com/responserms/spec/parser"

// VariablesDefinition allows registration of the vars definition block using
// NewWithSubset.
func VariablesDefinition() parser.NamedBlockDefinition {
	return &variablesDef{}
}

// EncryptionKeyDefinition allows registration of the encryption_key definition
// block using NewWithSubset.
func EncryptionKeyDefinition() parser.NamedBlockDefinition {
	return &encryptionKeyDef{}
}

// EventsDefinition allows registration of the events block using
// NewWithSubset.
func EventsDefinition() parser.NamedBlockDefinition {
	return &eventsDef{}
}

// DatabaseDefinition allows registration of the database block using
// NewWithSubset.
func DatabaseDefinition() parser.NamedBlockDefinition {
	return &databaseDef{}
}

func DeveloperDefinition() parser.NamedBlockDefinition {
	return &developerDef{}
}

func ClusterDefinition() parser.NamedBlockDefinition {
	return &clusterDef{}
}

// Schema defines the blocks and their respecify BlockDefinition by creating
// an instance of parser.BlockDefinitions.
func Schema() parser.NamedBlockDefinitions {
	return parser.NamedBlockDefinitions{

		// variables must be processed first
		VariablesDefinition(),

		// encryption_key block
		EncryptionKeyDefinition(),

		// // listener block
		// ListenerDefinition(),

		// // settings block
		// SettingsDefinition(),

		// database block
		DatabaseDefinition(),

		// pubsub block
		EventsDefinition(),

		// developer block
		DeveloperDefinition(),

		// cluster block
		ClusterDefinition(),
	}
}
