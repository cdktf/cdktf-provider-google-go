// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagecontrolfolderintelligenceconfig


type StorageControlFolderIntelligenceConfigFilter struct {
	// excluded_cloud_storage_buckets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/storage_control_folder_intelligence_config#excluded_cloud_storage_buckets StorageControlFolderIntelligenceConfig#excluded_cloud_storage_buckets}
	ExcludedCloudStorageBuckets *StorageControlFolderIntelligenceConfigFilterExcludedCloudStorageBuckets `field:"optional" json:"excludedCloudStorageBuckets" yaml:"excludedCloudStorageBuckets"`
	// excluded_cloud_storage_locations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/storage_control_folder_intelligence_config#excluded_cloud_storage_locations StorageControlFolderIntelligenceConfig#excluded_cloud_storage_locations}
	ExcludedCloudStorageLocations *StorageControlFolderIntelligenceConfigFilterExcludedCloudStorageLocations `field:"optional" json:"excludedCloudStorageLocations" yaml:"excludedCloudStorageLocations"`
	// included_cloud_storage_buckets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/storage_control_folder_intelligence_config#included_cloud_storage_buckets StorageControlFolderIntelligenceConfig#included_cloud_storage_buckets}
	IncludedCloudStorageBuckets *StorageControlFolderIntelligenceConfigFilterIncludedCloudStorageBuckets `field:"optional" json:"includedCloudStorageBuckets" yaml:"includedCloudStorageBuckets"`
	// included_cloud_storage_locations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/storage_control_folder_intelligence_config#included_cloud_storage_locations StorageControlFolderIntelligenceConfig#included_cloud_storage_locations}
	IncludedCloudStorageLocations *StorageControlFolderIntelligenceConfigFilterIncludedCloudStorageLocations `field:"optional" json:"includedCloudStorageLocations" yaml:"includedCloudStorageLocations"`
}

