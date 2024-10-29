// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageinsightsreportconfig


type StorageInsightsReportConfigFrequencyOptions struct {
	// end_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/storage_insights_report_config#end_date StorageInsightsReportConfig#end_date}
	EndDate *StorageInsightsReportConfigFrequencyOptionsEndDate `field:"required" json:"endDate" yaml:"endDate"`
	// The frequency in which inventory reports are generated. Values are DAILY or WEEKLY. Possible values: ["DAILY", "WEEKLY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/storage_insights_report_config#frequency StorageInsightsReportConfig#frequency}
	Frequency *string `field:"required" json:"frequency" yaml:"frequency"`
	// start_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/storage_insights_report_config#start_date StorageInsightsReportConfig#start_date}
	StartDate *StorageInsightsReportConfigFrequencyOptionsStartDate `field:"required" json:"startDate" yaml:"startDate"`
}

