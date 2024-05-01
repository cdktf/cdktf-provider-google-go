// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageinsightsreportconfig


type StorageInsightsReportConfigCsvOptions struct {
	// The delimiter used to separate the fields in the inventory report CSV file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/storage_insights_report_config#delimiter StorageInsightsReportConfig#delimiter}
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// The boolean that indicates whether or not headers are included in the inventory report CSV file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/storage_insights_report_config#header_required StorageInsightsReportConfig#header_required}
	HeaderRequired interface{} `field:"optional" json:"headerRequired" yaml:"headerRequired"`
	// The character used to separate the records in the inventory report CSV file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/storage_insights_report_config#record_separator StorageInsightsReportConfig#record_separator}
	RecordSeparator *string `field:"optional" json:"recordSeparator" yaml:"recordSeparator"`
}

