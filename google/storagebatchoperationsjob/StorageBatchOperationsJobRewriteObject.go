// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagebatchoperationsjob


type StorageBatchOperationsJobRewriteObject struct {
	// valid kms key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/storage_batch_operations_job#kms_key StorageBatchOperationsJob#kms_key}
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
}

