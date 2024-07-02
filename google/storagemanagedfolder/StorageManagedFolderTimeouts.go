// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagemanagedfolder


type StorageManagedFolderTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.36.0/docs/resources/storage_managed_folder#create StorageManagedFolder#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.36.0/docs/resources/storage_managed_folder#delete StorageManagedFolder#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

