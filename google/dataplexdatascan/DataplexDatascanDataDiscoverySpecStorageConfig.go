// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexdatascan


type DataplexDatascanDataDiscoverySpecStorageConfig struct {
	// csv_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dataplex_datascan#csv_options DataplexDatascan#csv_options}
	CsvOptions *DataplexDatascanDataDiscoverySpecStorageConfigCsvOptions `field:"optional" json:"csvOptions" yaml:"csvOptions"`
	// Defines the data to exclude during discovery.
	//
	// Provide a list of patterns that identify the data to exclude. For Cloud Storage bucket assets, these patterns are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these patterns are interpreted as patterns to match table names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dataplex_datascan#exclude_patterns DataplexDatascan#exclude_patterns}
	ExcludePatterns *[]*string `field:"optional" json:"excludePatterns" yaml:"excludePatterns"`
	// Defines the data to include during discovery when only a subset of the data should be considered.
	//
	// Provide a list of patterns that identify the data to include. For Cloud Storage bucket assets, these patterns are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these patterns are interpreted as patterns to match table names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dataplex_datascan#include_patterns DataplexDatascan#include_patterns}
	IncludePatterns *[]*string `field:"optional" json:"includePatterns" yaml:"includePatterns"`
	// json_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dataplex_datascan#json_options DataplexDatascan#json_options}
	JsonOptions *DataplexDatascanDataDiscoverySpecStorageConfigJsonOptions `field:"optional" json:"jsonOptions" yaml:"jsonOptions"`
}

