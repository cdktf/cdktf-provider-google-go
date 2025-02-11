// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iaptunneldestgroup


type IapTunnelDestGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/iap_tunnel_dest_group#create IapTunnelDestGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/iap_tunnel_dest_group#delete IapTunnelDestGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/iap_tunnel_dest_group#update IapTunnelDestGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

