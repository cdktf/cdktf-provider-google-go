// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageinsightsreportconfig


type StorageInsightsReportConfigObjectMetadataReportOptionsStorageFilters struct {
	// The filter to use when specifying which bucket to generate inventory reports for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/storage_insights_report_config#bucket StorageInsightsReportConfig#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
}

