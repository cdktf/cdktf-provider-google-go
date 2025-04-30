// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iaptunnelinstanceiammember


type IapTunnelInstanceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/iap_tunnel_instance_iam_member#expression IapTunnelInstanceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/iap_tunnel_instance_iam_member#title IapTunnelInstanceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/iap_tunnel_instance_iam_member#description IapTunnelInstanceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

