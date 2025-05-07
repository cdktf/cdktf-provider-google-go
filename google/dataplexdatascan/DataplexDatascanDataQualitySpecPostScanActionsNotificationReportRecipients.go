// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexdatascan


type DataplexDatascanDataQualitySpecPostScanActionsNotificationReportRecipients struct {
	// The email recipients who will receive the DataQualityScan results report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/dataplex_datascan#emails DataplexDatascan#emails}
	Emails *[]*string `field:"optional" json:"emails" yaml:"emails"`
}

