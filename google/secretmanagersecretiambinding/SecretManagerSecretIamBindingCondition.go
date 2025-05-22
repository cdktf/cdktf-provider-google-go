// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secretmanagersecretiambinding


type SecretManagerSecretIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/secret_manager_secret_iam_binding#expression SecretManagerSecretIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/secret_manager_secret_iam_binding#title SecretManagerSecretIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/secret_manager_secret_iam_binding#description SecretManagerSecretIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

