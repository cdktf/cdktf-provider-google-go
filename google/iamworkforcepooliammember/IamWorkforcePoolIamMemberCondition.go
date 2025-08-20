// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamworkforcepooliammember


type IamWorkforcePoolIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/iam_workforce_pool_iam_member#expression IamWorkforcePoolIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/iam_workforce_pool_iam_member#title IamWorkforcePoolIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/iam_workforce_pool_iam_member#description IamWorkforcePoolIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

