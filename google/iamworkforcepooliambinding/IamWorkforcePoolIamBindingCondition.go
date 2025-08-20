// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamworkforcepooliambinding


type IamWorkforcePoolIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/iam_workforce_pool_iam_binding#expression IamWorkforcePoolIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/iam_workforce_pool_iam_binding#title IamWorkforcePoolIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/iam_workforce_pool_iam_binding#description IamWorkforcePoolIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

