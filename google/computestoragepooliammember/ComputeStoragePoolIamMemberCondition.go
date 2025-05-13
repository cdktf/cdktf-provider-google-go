// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computestoragepooliammember


type ComputeStoragePoolIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/compute_storage_pool_iam_member#expression ComputeStoragePoolIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/compute_storage_pool_iam_member#title ComputeStoragePoolIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/compute_storage_pool_iam_member#description ComputeStoragePoolIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

