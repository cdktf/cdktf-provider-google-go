// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexdatascan


type DataplexDatascanDataDiscoverySpec struct {
	// bigquery_publishing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dataplex_datascan#bigquery_publishing_config DataplexDatascan#bigquery_publishing_config}
	BigqueryPublishingConfig *DataplexDatascanDataDiscoverySpecBigqueryPublishingConfig `field:"optional" json:"bigqueryPublishingConfig" yaml:"bigqueryPublishingConfig"`
	// storage_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dataplex_datascan#storage_config DataplexDatascan#storage_config}
	StorageConfig *DataplexDatascanDataDiscoverySpecStorageConfig `field:"optional" json:"storageConfig" yaml:"storageConfig"`
}

