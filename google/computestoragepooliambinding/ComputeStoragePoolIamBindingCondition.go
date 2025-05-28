// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computestoragepooliambinding


type ComputeStoragePoolIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/compute_storage_pool_iam_binding#expression ComputeStoragePoolIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/compute_storage_pool_iam_binding#title ComputeStoragePoolIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/compute_storage_pool_iam_binding#description ComputeStoragePoolIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

