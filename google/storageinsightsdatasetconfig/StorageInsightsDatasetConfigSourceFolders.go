// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageinsightsdatasetconfig


type StorageInsightsDatasetConfigSourceFolders struct {
	// The list of folder numbers to include in the DatasetConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/storage_insights_dataset_config#folder_numbers StorageInsightsDatasetConfig#folder_numbers}
	FolderNumbers *[]*string `field:"optional" json:"folderNumbers" yaml:"folderNumbers"`
}

