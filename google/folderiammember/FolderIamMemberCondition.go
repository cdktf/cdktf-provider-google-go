// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package folderiammember


type FolderIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/folder_iam_member#expression FolderIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/folder_iam_member#title FolderIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/folder_iam_member#description FolderIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

