// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package geminirepositorygroupiambinding


type GeminiRepositoryGroupIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/gemini_repository_group_iam_binding#expression GeminiRepositoryGroupIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/gemini_repository_group_iam_binding#title GeminiRepositoryGroupIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/gemini_repository_group_iam_binding#description GeminiRepositoryGroupIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

