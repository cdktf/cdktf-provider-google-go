// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagemanagedfolderiambinding


type StorageManagedFolderIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/storage_managed_folder_iam_binding#expression StorageManagedFolderIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/storage_managed_folder_iam_binding#title StorageManagedFolderIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/storage_managed_folder_iam_binding#description StorageManagedFolderIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

