// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notebooksruntimeiammember


type NotebooksRuntimeIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/notebooks_runtime_iam_member#expression NotebooksRuntimeIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/notebooks_runtime_iam_member#title NotebooksRuntimeIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/notebooks_runtime_iam_member#description NotebooksRuntimeIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

