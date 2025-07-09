// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexdatascan


type DataplexDatascanDataQualitySpecPostScanActionsNotificationReport struct {
	// recipients block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/dataplex_datascan#recipients DataplexDatascan#recipients}
	Recipients *DataplexDatascanDataQualitySpecPostScanActionsNotificationReportRecipients `field:"required" json:"recipients" yaml:"recipients"`
	// job_end_trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/dataplex_datascan#job_end_trigger DataplexDatascan#job_end_trigger}
	JobEndTrigger *DataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobEndTrigger `field:"optional" json:"jobEndTrigger" yaml:"jobEndTrigger"`
	// job_failure_trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/dataplex_datascan#job_failure_trigger DataplexDatascan#job_failure_trigger}
	JobFailureTrigger *DataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobFailureTrigger `field:"optional" json:"jobFailureTrigger" yaml:"jobFailureTrigger"`
	// score_threshold_trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/dataplex_datascan#score_threshold_trigger DataplexDatascan#score_threshold_trigger}
	ScoreThresholdTrigger *DataplexDatascanDataQualitySpecPostScanActionsNotificationReportScoreThresholdTrigger `field:"optional" json:"scoreThresholdTrigger" yaml:"scoreThresholdTrigger"`
}

