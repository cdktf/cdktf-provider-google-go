// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageinsightsreportconfig


type StorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptions struct {
	// The destination bucket that stores the generated inventory reports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/storage_insights_report_config#bucket StorageInsightsReportConfig#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The path within the destination bucket to store generated inventory reports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/storage_insights_report_config#destination_path StorageInsightsReportConfig#destination_path}
	DestinationPath *string `field:"optional" json:"destinationPath" yaml:"destinationPath"`
}

