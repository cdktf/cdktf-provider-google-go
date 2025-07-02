// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudidentitygroupmembership


type CloudIdentityGroupMembershipTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/cloud_identity_group_membership#create CloudIdentityGroupMembership#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/cloud_identity_group_membership#delete CloudIdentityGroupMembership#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/cloud_identity_group_membership#update CloudIdentityGroupMembership#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

