// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sourcereporepositoryiambinding


type SourcerepoRepositoryIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/sourcerepo_repository_iam_binding#expression SourcerepoRepositoryIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/sourcerepo_repository_iam_binding#title SourcerepoRepositoryIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/sourcerepo_repository_iam_binding#description SourcerepoRepositoryIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

