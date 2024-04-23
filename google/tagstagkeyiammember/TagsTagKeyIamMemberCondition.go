// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tagstagkeyiammember


type TagsTagKeyIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.26.0/docs/resources/tags_tag_key_iam_member#expression TagsTagKeyIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.26.0/docs/resources/tags_tag_key_iam_member#title TagsTagKeyIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.26.0/docs/resources/tags_tag_key_iam_member#description TagsTagKeyIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

