// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagebatchoperationsjob


type StorageBatchOperationsJobBucketListStruct struct {
	// buckets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/storage_batch_operations_job#buckets StorageBatchOperationsJob#buckets}
	Buckets *StorageBatchOperationsJobBucketListBuckets `field:"required" json:"buckets" yaml:"buckets"`
}

