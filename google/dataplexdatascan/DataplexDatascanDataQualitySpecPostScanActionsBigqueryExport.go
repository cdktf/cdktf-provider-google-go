// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexdatascan


type DataplexDatascanDataQualitySpecPostScanActionsBigqueryExport struct {
	// The BigQuery table to export DataQualityScan results to. Format://bigquery.googleapis.com/projects/PROJECT_ID/datasets/DATASET_ID/tables/TABLE_ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.35.0/docs/resources/dataplex_datascan#results_table DataplexDatascan#results_table}
	ResultsTable *string `field:"optional" json:"resultsTable" yaml:"resultsTable"`
}

