// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexdatascan


type DataplexDatascanDataQualitySpec struct {
	// post_scan_actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/dataplex_datascan#post_scan_actions DataplexDatascan#post_scan_actions}
	PostScanActions *DataplexDatascanDataQualitySpecPostScanActions `field:"optional" json:"postScanActions" yaml:"postScanActions"`
	// A filter applied to all rows in a single DataScan job.
	//
	// The filter needs to be a valid SQL expression for a WHERE clause in BigQuery standard SQL syntax. Example: col1 >= 0 AND col2 < 10
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/dataplex_datascan#row_filter DataplexDatascan#row_filter}
	RowFilter *string `field:"optional" json:"rowFilter" yaml:"rowFilter"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/dataplex_datascan#rules DataplexDatascan#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// The percentage of the records to be selected from the dataset for DataScan.
	//
	// Value can range between 0.0 and 100.0 with up to 3 significant decimal digits.
	// Sampling is not applied if 'sampling_percent' is not specified, 0 or 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/dataplex_datascan#sampling_percent DataplexDatascan#sampling_percent}
	SamplingPercent *float64 `field:"optional" json:"samplingPercent" yaml:"samplingPercent"`
}

