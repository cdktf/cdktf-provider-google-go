// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgenetworksubnet


type EdgenetworkSubnetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.16.0/docs/resources/edgenetwork_subnet#create EdgenetworkSubnet#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.16.0/docs/resources/edgenetwork_subnet#delete EdgenetworkSubnet#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

