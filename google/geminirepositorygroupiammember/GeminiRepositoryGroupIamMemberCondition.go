// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package geminirepositorygroupiammember


type GeminiRepositoryGroupIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/gemini_repository_group_iam_member#expression GeminiRepositoryGroupIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/gemini_repository_group_iam_member#title GeminiRepositoryGroupIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/gemini_repository_group_iam_member#description GeminiRepositoryGroupIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

