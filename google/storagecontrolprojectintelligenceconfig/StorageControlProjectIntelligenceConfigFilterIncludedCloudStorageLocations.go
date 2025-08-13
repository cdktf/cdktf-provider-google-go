// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagecontrolprojectintelligenceconfig


type StorageControlProjectIntelligenceConfigFilterIncludedCloudStorageLocations struct {
	// List of locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/storage_control_project_intelligence_config#locations StorageControlProjectIntelligenceConfig#locations}
	Locations *[]*string `field:"required" json:"locations" yaml:"locations"`
}

