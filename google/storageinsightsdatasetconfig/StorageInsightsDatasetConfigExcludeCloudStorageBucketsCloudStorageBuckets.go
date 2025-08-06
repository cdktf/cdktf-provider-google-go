// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageinsightsdatasetconfig


type StorageInsightsDatasetConfigExcludeCloudStorageBucketsCloudStorageBuckets struct {
	// The list of cloud storage bucket names to exclude in the DatasetConfig.
	//
	// Exactly one of the bucket_name and bucket_prefix_regex should be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/storage_insights_dataset_config#bucket_name StorageInsightsDatasetConfig#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The list of regex patterns for bucket names matching the regex.
	//
	// Regex should follow the syntax specified in google/re2 on GitHub.
	// Exactly one of the bucket_name and bucket_prefix_regex should be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/storage_insights_dataset_config#bucket_prefix_regex StorageInsightsDatasetConfig#bucket_prefix_regex}
	BucketPrefixRegex *string `field:"optional" json:"bucketPrefixRegex" yaml:"bucketPrefixRegex"`
}

