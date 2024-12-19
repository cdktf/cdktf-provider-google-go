// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secretmanagerregionalsecretiammember


type SecretManagerRegionalSecretIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/secret_manager_regional_secret_iam_member#expression SecretManagerRegionalSecretIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/secret_manager_regional_secret_iam_member#title SecretManagerRegionalSecretIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/secret_manager_regional_secret_iam_member#description SecretManagerRegionalSecretIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

