// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securesourcemanagerrepositoryiammember


type SecureSourceManagerRepositoryIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/secure_source_manager_repository_iam_member#expression SecureSourceManagerRepositoryIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/secure_source_manager_repository_iam_member#title SecureSourceManagerRepositoryIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/secure_source_manager_repository_iam_member#description SecureSourceManagerRepositoryIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

