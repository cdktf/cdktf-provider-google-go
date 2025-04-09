// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeeenvironmentiammember


type ApigeeEnvironmentIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/apigee_environment_iam_member#expression ApigeeEnvironmentIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/apigee_environment_iam_member#title ApigeeEnvironmentIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/apigee_environment_iam_member#description ApigeeEnvironmentIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

