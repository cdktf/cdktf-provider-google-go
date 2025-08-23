// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagebatchoperationsjob


type StorageBatchOperationsJobBucketListBuckets struct {
	// Bucket name for the objects to be transformed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/storage_batch_operations_job#bucket StorageBatchOperationsJob#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// manifest block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/storage_batch_operations_job#manifest StorageBatchOperationsJob#manifest}
	Manifest *StorageBatchOperationsJobBucketListBucketsManifest `field:"optional" json:"manifest" yaml:"manifest"`
	// prefix_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/storage_batch_operations_job#prefix_list StorageBatchOperationsJob#prefix_list}
	PrefixList *StorageBatchOperationsJobBucketListBucketsPrefixListStruct `field:"optional" json:"prefixList" yaml:"prefixList"`
}

