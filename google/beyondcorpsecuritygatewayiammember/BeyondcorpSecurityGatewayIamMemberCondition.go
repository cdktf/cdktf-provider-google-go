// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package beyondcorpsecuritygatewayiammember


type BeyondcorpSecurityGatewayIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/beyondcorp_security_gateway_iam_member#expression BeyondcorpSecurityGatewayIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/beyondcorp_security_gateway_iam_member#title BeyondcorpSecurityGatewayIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/beyondcorp_security_gateway_iam_member#description BeyondcorpSecurityGatewayIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

