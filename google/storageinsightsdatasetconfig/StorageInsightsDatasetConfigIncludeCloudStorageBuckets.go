// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageinsightsdatasetconfig


type StorageInsightsDatasetConfigIncludeCloudStorageBuckets struct {
	// cloud_storage_buckets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/storage_insights_dataset_config#cloud_storage_buckets StorageInsightsDatasetConfig#cloud_storage_buckets}
	CloudStorageBuckets interface{} `field:"required" json:"cloudStorageBuckets" yaml:"cloudStorageBuckets"`
}

