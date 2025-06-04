// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package beyondcorpsecuritygatewayapplicationiammember


type BeyondcorpSecurityGatewayApplicationIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/beyondcorp_security_gateway_application_iam_member#expression BeyondcorpSecurityGatewayApplicationIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/beyondcorp_security_gateway_application_iam_member#title BeyondcorpSecurityGatewayApplicationIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/beyondcorp_security_gateway_application_iam_member#description BeyondcorpSecurityGatewayApplicationIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

