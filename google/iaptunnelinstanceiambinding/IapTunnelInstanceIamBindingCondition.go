// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iaptunnelinstanceiambinding


type IapTunnelInstanceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/iap_tunnel_instance_iam_binding#expression IapTunnelInstanceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/iap_tunnel_instance_iam_binding#title IapTunnelInstanceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/iap_tunnel_instance_iam_binding#description IapTunnelInstanceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

