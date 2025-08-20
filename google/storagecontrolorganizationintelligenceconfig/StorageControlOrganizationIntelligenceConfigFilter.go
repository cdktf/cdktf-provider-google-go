// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagecontrolorganizationintelligenceconfig


type StorageControlOrganizationIntelligenceConfigFilter struct {
	// excluded_cloud_storage_buckets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/storage_control_organization_intelligence_config#excluded_cloud_storage_buckets StorageControlOrganizationIntelligenceConfig#excluded_cloud_storage_buckets}
	ExcludedCloudStorageBuckets *StorageControlOrganizationIntelligenceConfigFilterExcludedCloudStorageBuckets `field:"optional" json:"excludedCloudStorageBuckets" yaml:"excludedCloudStorageBuckets"`
	// excluded_cloud_storage_locations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/storage_control_organization_intelligence_config#excluded_cloud_storage_locations StorageControlOrganizationIntelligenceConfig#excluded_cloud_storage_locations}
	ExcludedCloudStorageLocations *StorageControlOrganizationIntelligenceConfigFilterExcludedCloudStorageLocations `field:"optional" json:"excludedCloudStorageLocations" yaml:"excludedCloudStorageLocations"`
	// included_cloud_storage_buckets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/storage_control_organization_intelligence_config#included_cloud_storage_buckets StorageControlOrganizationIntelligenceConfig#included_cloud_storage_buckets}
	IncludedCloudStorageBuckets *StorageControlOrganizationIntelligenceConfigFilterIncludedCloudStorageBuckets `field:"optional" json:"includedCloudStorageBuckets" yaml:"includedCloudStorageBuckets"`
	// included_cloud_storage_locations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/storage_control_organization_intelligence_config#included_cloud_storage_locations StorageControlOrganizationIntelligenceConfig#included_cloud_storage_locations}
	IncludedCloudStorageLocations *StorageControlOrganizationIntelligenceConfigFilterIncludedCloudStorageLocations `field:"optional" json:"includedCloudStorageLocations" yaml:"includedCloudStorageLocations"`
}

