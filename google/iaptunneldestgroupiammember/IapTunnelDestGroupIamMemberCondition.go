// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iaptunneldestgroupiammember


type IapTunnelDestGroupIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/iap_tunnel_dest_group_iam_member#expression IapTunnelDestGroupIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/iap_tunnel_dest_group_iam_member#title IapTunnelDestGroupIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/iap_tunnel_dest_group_iam_member#description IapTunnelDestGroupIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

