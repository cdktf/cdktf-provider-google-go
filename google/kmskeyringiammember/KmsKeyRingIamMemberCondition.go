// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kmskeyringiammember


type KmsKeyRingIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/kms_key_ring_iam_member#expression KmsKeyRingIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/kms_key_ring_iam_member#title KmsKeyRingIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/kms_key_ring_iam_member#description KmsKeyRingIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

