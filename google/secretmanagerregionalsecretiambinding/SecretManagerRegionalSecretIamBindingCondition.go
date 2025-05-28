// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secretmanagerregionalsecretiambinding


type SecretManagerRegionalSecretIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/secret_manager_regional_secret_iam_binding#expression SecretManagerRegionalSecretIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/secret_manager_regional_secret_iam_binding#title SecretManagerRegionalSecretIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/secret_manager_regional_secret_iam_binding#description SecretManagerRegionalSecretIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

