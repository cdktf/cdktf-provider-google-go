// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgecontainervpnconnection


type EdgecontainerVpnConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.1.0/docs/resources/edgecontainer_vpn_connection#create EdgecontainerVpnConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.1.0/docs/resources/edgecontainer_vpn_connection#delete EdgecontainerVpnConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.1.0/docs/resources/edgecontainer_vpn_connection#update EdgecontainerVpnConnection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

