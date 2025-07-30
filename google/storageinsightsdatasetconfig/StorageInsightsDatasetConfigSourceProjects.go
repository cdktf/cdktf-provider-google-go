// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageinsightsdatasetconfig


type StorageInsightsDatasetConfigSourceProjects struct {
	// The list of project numbers to include in the DatasetConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/storage_insights_dataset_config#project_numbers StorageInsightsDatasetConfig#project_numbers}
	ProjectNumbers *[]*string `field:"optional" json:"projectNumbers" yaml:"projectNumbers"`
}

