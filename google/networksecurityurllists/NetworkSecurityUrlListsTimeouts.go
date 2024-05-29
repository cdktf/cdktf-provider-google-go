// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityurllists


type NetworkSecurityUrlListsTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/network_security_url_lists#create NetworkSecurityUrlLists#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/network_security_url_lists#delete NetworkSecurityUrlLists#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/network_security_url_lists#update NetworkSecurityUrlLists#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

