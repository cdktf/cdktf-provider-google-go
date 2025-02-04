// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagetransferjob


type StorageTransferJobTransferSpecHdfsDataSource struct {
	// Directory path to the filesystem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/storage_transfer_job#path StorageTransferJob#path}
	Path *string `field:"required" json:"path" yaml:"path"`
}

