// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageinsightsdatasetconfig


type StorageInsightsDatasetConfigIncludeCloudStorageLocations struct {
	// The list of cloud storage locations to include in the DatasetConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/storage_insights_dataset_config#locations StorageInsightsDatasetConfig#locations}
	Locations *[]*string `field:"required" json:"locations" yaml:"locations"`
}

