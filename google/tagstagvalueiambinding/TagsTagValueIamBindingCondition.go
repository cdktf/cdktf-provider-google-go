// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tagstagvalueiambinding


type TagsTagValueIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.20.0/docs/resources/tags_tag_value_iam_binding#expression TagsTagValueIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.20.0/docs/resources/tags_tag_value_iam_binding#title TagsTagValueIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.20.0/docs/resources/tags_tag_value_iam_binding#description TagsTagValueIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

