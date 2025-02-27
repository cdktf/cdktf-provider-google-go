// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iaptunneliambinding


type IapTunnelIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/iap_tunnel_iam_binding#expression IapTunnelIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/iap_tunnel_iam_binding#title IapTunnelIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/iap_tunnel_iam_binding#description IapTunnelIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

