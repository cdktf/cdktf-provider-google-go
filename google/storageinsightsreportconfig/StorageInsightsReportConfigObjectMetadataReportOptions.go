// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageinsightsreportconfig


type StorageInsightsReportConfigObjectMetadataReportOptions struct {
	// The metadata fields included in an inventory report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/storage_insights_report_config#metadata_fields StorageInsightsReportConfig#metadata_fields}
	MetadataFields *[]*string `field:"required" json:"metadataFields" yaml:"metadataFields"`
	// storage_destination_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/storage_insights_report_config#storage_destination_options StorageInsightsReportConfig#storage_destination_options}
	StorageDestinationOptions *StorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptions `field:"required" json:"storageDestinationOptions" yaml:"storageDestinationOptions"`
	// storage_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/storage_insights_report_config#storage_filters StorageInsightsReportConfig#storage_filters}
	StorageFilters *StorageInsightsReportConfigObjectMetadataReportOptionsStorageFilters `field:"optional" json:"storageFilters" yaml:"storageFilters"`
}

