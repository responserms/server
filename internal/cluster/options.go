// Copyright (c) 2021 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cluster

import (
	"fmt"

	"github.com/buraksezer/olric/config"
	olricnats "github.com/justinfx/olric-nats-plugin/lib"
)

type AutoJoinOptions struct {
	Type    string
	URL     string
	Subject string
}

// Options configures the Cluster.
type Options struct {
	EnvType           string
	BindAddress       string
	BindPort          int
	MemberBindAddress string
	MemberBindPort    int
	AutoJoin          *AutoJoinOptions
	olric             *config.Config
}

// init initializes an Olric config from Options to initialize the cluster.
func (o *Options) init() error {
	olric := config.New(o.EnvType)

	olric.BindAddr = o.BindAddress
	olric.BindPort = o.BindPort

	mlist, err := config.NewMemberlistConfig(o.EnvType)
	if err != nil {
		return fmt.Errorf("new mlist: %w", err)
	}

	mlist.BindAddr = o.MemberBindAddress
	mlist.BindPort = o.MemberBindPort

	mlist.AdvertiseAddr = o.MemberBindAddress
	mlist.AdvertisePort = o.MemberBindPort

	olric.MemberlistConfig = mlist

	o.olric = olric

	if o.AutoJoin != nil {
		if o.AutoJoin.Type == "nats" {
			if err := o.configureNATS(); err != nil {
				return fmt.Errorf("configure nats autojoin: %w", err)
			}
		}
	}

	return nil
}

func (o *Options) configureNATS() error {
	sd := make(map[string]interface{})
	sd["plugin"] = &olricnats.NatsDiscovery{
		Config: &olricnats.Config{
			Provider: "nats",
			Url:      o.AutoJoin.URL,
			Subject:  o.AutoJoin.Subject,
			Payload: olricnats.Payload{
				Host: o.MemberBindAddress,
				Port: o.MemberBindPort,
			},
		},
	}

	o.olric.ServiceDiscovery = sd

	return nil
}
