// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagemanagedfolderiammember


type StorageManagedFolderIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/storage_managed_folder_iam_member#expression StorageManagedFolderIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/storage_managed_folder_iam_member#title StorageManagedFolderIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/storage_managed_folder_iam_member#description StorageManagedFolderIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

