// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexdatascan


type DataplexDatascanDataQualitySpecPostScanActions struct {
	// bigquery_export block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/dataplex_datascan#bigquery_export DataplexDatascan#bigquery_export}
	BigqueryExport *DataplexDatascanDataQualitySpecPostScanActionsBigqueryExport `field:"optional" json:"bigqueryExport" yaml:"bigqueryExport"`
}

