// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeployautomation


type ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetry struct {
	// Required.
	//
	// Total number of retries. Retry is skipped if set to 0; The minimum value is 1, and the maximum value is 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/clouddeploy_automation#attempts ClouddeployAutomation#attempts}
	Attempts *string `field:"required" json:"attempts" yaml:"attempts"`
	// Optional.
	//
	// The pattern of how wait time will be increased. Default is linear. Backoff mode will be ignored if wait is 0. Possible values: ["BACKOFF_MODE_UNSPECIFIED", "BACKOFF_MODE_LINEAR", "BACKOFF_MODE_EXPONENTIAL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/clouddeploy_automation#backoff_mode ClouddeployAutomation#backoff_mode}
	BackoffMode *string `field:"optional" json:"backoffMode" yaml:"backoffMode"`
	// Optional.
	//
	// How long to wait for the first retry. Default is 0, and the maximum value is 14d. A duration in seconds with up to nine fractional digits, ending with 's'. Example: '3.5s'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/clouddeploy_automation#wait ClouddeployAutomation#wait}
	Wait *string `field:"optional" json:"wait" yaml:"wait"`
}

