// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package schema

import "github.com/hashicorp/go-sockaddr"

// ensure interfaces are implemented properly
var _ PublicIPGetter = (*IPGetter)(nil)
var _ PrivateIPGetter = (*IPGetter)(nil)
var _ InterfaceIPGetter = (*IPGetter)(nil)

// PrivateIPGetter implements a function that returns the current private IP address for the
// runtime environment.
type PrivateIPGetter interface {
	GetPrivateIP() (string, error)
}

// PublicIPGetter implements a function that returns the current public IP address for the
// runtime environment.
type PublicIPGetter interface {
	GetPublicIP() (string, error)
}

// InterfaceIPGetter implements a function that returns the current IP address for the named
// interface provided.
type InterfaceIPGetter interface {
	GetInterfaceIP(namedIf string) (string, error)
}

// IPGetter implements the PrivateIPGetter, PublicIPGetter, and InterfaceIPGetter interfaces
// used by the networking functions to return proper addresses for listener bindings.
type IPGetter struct{}

// GetPrivateIP returns a string with a single IP address that is part of RFC 6890 and has a
// default route. If the system can't determine its IP address or find an RFC 6890 IP address,
// an empty string will be returned instead.
func (g *IPGetter) GetPrivateIP() (string, error) {
	return sockaddr.GetPrivateIP()
}

// GetPublicIP returns a string with a single IP address that is NOT part of RFC 6890 and has a
// default route. If the system can't determine its IP address or find a non RFC 6890 IP address,
// an empty string will be returned instead.
func (g *IPGetter) GetPublicIP() (string, error) {
	return sockaddr.GetPublicIP()
}

// GetInterfaceIP returns a string with a single IP address sorted by the size of the network
// (i.e. IP addresses with a smaller netmask, larger network size, are sorted first). This function is
// the `eval` equivalent of:
func (g *IPGetter) GetInterfaceIP(namedIf string) (string, error) {
	return sockaddr.GetInterfaceIP(namedIf)
}
