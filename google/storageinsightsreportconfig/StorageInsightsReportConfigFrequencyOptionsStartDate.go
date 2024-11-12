// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageinsightsreportconfig


type StorageInsightsReportConfigFrequencyOptionsStartDate struct {
	// The day of the month to start generating inventory reports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/storage_insights_report_config#day StorageInsightsReportConfig#day}
	Day *float64 `field:"required" json:"day" yaml:"day"`
	// The month to start generating inventory reports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/storage_insights_report_config#month StorageInsightsReportConfig#month}
	Month *float64 `field:"required" json:"month" yaml:"month"`
	// The year to start generating inventory reports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/storage_insights_report_config#year StorageInsightsReportConfig#year}
	Year *float64 `field:"required" json:"year" yaml:"year"`
}

