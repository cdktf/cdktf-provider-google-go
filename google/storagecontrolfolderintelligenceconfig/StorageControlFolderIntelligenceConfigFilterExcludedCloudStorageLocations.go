// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagecontrolfolderintelligenceconfig


type StorageControlFolderIntelligenceConfigFilterExcludedCloudStorageLocations struct {
	// List of locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/storage_control_folder_intelligence_config#locations StorageControlFolderIntelligenceConfig#locations}
	Locations *[]*string `field:"required" json:"locations" yaml:"locations"`
}

