// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgenetworksubnet


type EdgenetworkSubnetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/edgenetwork_subnet#create EdgenetworkSubnet#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/edgenetwork_subnet#delete EdgenetworkSubnet#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/edgenetwork_subnet#update EdgenetworkSubnet#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

