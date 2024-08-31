// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package binaryauthorizationattestoriambinding


type BinaryAuthorizationAttestorIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/binary_authorization_attestor_iam_binding#expression BinaryAuthorizationAttestorIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/binary_authorization_attestor_iam_binding#title BinaryAuthorizationAttestorIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/binary_authorization_attestor_iam_binding#description BinaryAuthorizationAttestorIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

