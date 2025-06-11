// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagetransferjob


type StorageTransferJobTransferSpecTransferOptionsMetadataOptions struct {
	// Specifies how each object's ACLs should be preserved for transfers between Google Cloud Storage buckets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/storage_transfer_job#acl StorageTransferJob#acl}
	Acl *string `field:"optional" json:"acl" yaml:"acl"`
	// Specifies how each file's POSIX group ID (GID) attribute should be handled by the transfer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/storage_transfer_job#gid StorageTransferJob#gid}
	Gid *string `field:"optional" json:"gid" yaml:"gid"`
	// Specifies how each object's Cloud KMS customer-managed encryption key (CMEK) is preserved for transfers between Google Cloud Storage buckets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/storage_transfer_job#kms_key StorageTransferJob#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Specifies how each file's mode attribute should be handled by the transfer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/storage_transfer_job#mode StorageTransferJob#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Specifies the storage class to set on objects being transferred to Google Cloud Storage buckets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/storage_transfer_job#storage_class StorageTransferJob#storage_class}
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
	// Specifies how symlinks should be handled by the transfer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/storage_transfer_job#symlink StorageTransferJob#symlink}
	Symlink *string `field:"optional" json:"symlink" yaml:"symlink"`
	// SSpecifies how each object's temporary hold status should be preserved for transfers between Google Cloud Storage buckets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/storage_transfer_job#temporary_hold StorageTransferJob#temporary_hold}
	TemporaryHold *string `field:"optional" json:"temporaryHold" yaml:"temporaryHold"`
	// Specifies how each object's timeCreated metadata is preserved for transfers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/storage_transfer_job#time_created StorageTransferJob#time_created}
	TimeCreated *string `field:"optional" json:"timeCreated" yaml:"timeCreated"`
	// Specifies how each file's POSIX user ID (UID) attribute should be handled by the transfer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/storage_transfer_job#uid StorageTransferJob#uid}
	Uid *string `field:"optional" json:"uid" yaml:"uid"`
}

