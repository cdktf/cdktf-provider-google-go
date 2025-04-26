// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagecontrolprojectintelligenceconfig


type StorageControlProjectIntelligenceConfigFilter struct {
	// excluded_cloud_storage_buckets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/storage_control_project_intelligence_config#excluded_cloud_storage_buckets StorageControlProjectIntelligenceConfig#excluded_cloud_storage_buckets}
	ExcludedCloudStorageBuckets *StorageControlProjectIntelligenceConfigFilterExcludedCloudStorageBuckets `field:"optional" json:"excludedCloudStorageBuckets" yaml:"excludedCloudStorageBuckets"`
	// excluded_cloud_storage_locations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/storage_control_project_intelligence_config#excluded_cloud_storage_locations StorageControlProjectIntelligenceConfig#excluded_cloud_storage_locations}
	ExcludedCloudStorageLocations *StorageControlProjectIntelligenceConfigFilterExcludedCloudStorageLocations `field:"optional" json:"excludedCloudStorageLocations" yaml:"excludedCloudStorageLocations"`
	// included_cloud_storage_buckets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/storage_control_project_intelligence_config#included_cloud_storage_buckets StorageControlProjectIntelligenceConfig#included_cloud_storage_buckets}
	IncludedCloudStorageBuckets *StorageControlProjectIntelligenceConfigFilterIncludedCloudStorageBuckets `field:"optional" json:"includedCloudStorageBuckets" yaml:"includedCloudStorageBuckets"`
	// included_cloud_storage_locations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/storage_control_project_intelligence_config#included_cloud_storage_locations StorageControlProjectIntelligenceConfig#included_cloud_storage_locations}
	IncludedCloudStorageLocations *StorageControlProjectIntelligenceConfigFilterIncludedCloudStorageLocations `field:"optional" json:"includedCloudStorageLocations" yaml:"includedCloudStorageLocations"`
}

