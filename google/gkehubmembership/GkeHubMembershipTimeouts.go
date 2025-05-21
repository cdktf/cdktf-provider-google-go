// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubmembership


type GkeHubMembershipTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/gke_hub_membership#create GkeHubMembership#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/gke_hub_membership#delete GkeHubMembership#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/gke_hub_membership#update GkeHubMembership#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

