// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package beyondcorpsecuritygatewayiambinding


type BeyondcorpSecurityGatewayIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/beyondcorp_security_gateway_iam_binding#expression BeyondcorpSecurityGatewayIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/beyondcorp_security_gateway_iam_binding#title BeyondcorpSecurityGatewayIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/beyondcorp_security_gateway_iam_binding#description BeyondcorpSecurityGatewayIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

