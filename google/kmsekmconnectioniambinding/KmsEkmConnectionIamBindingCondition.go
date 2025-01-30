// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kmsekmconnectioniambinding


type KmsEkmConnectionIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/kms_ekm_connection_iam_binding#expression KmsEkmConnectionIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/kms_ekm_connection_iam_binding#title KmsEkmConnectionIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/kms_ekm_connection_iam_binding#description KmsEkmConnectionIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

