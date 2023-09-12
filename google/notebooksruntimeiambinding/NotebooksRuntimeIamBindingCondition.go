// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notebooksruntimeiambinding


type NotebooksRuntimeIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.82.0/docs/resources/notebooks_runtime_iam_binding#expression NotebooksRuntimeIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.82.0/docs/resources/notebooks_runtime_iam_binding#title NotebooksRuntimeIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.82.0/docs/resources/notebooks_runtime_iam_binding#description NotebooksRuntimeIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

