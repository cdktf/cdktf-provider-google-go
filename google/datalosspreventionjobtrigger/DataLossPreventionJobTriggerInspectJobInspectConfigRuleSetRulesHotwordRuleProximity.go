// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity struct {
	// Number of characters after the finding to consider. Either this or window_before must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/data_loss_prevention_job_trigger#window_after DataLossPreventionJobTrigger#window_after}
	WindowAfter *float64 `field:"optional" json:"windowAfter" yaml:"windowAfter"`
	// Number of characters before the finding to consider. Either this or window_after must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/data_loss_prevention_job_trigger#window_before DataLossPreventionJobTrigger#window_before}
	WindowBefore *float64 `field:"optional" json:"windowBefore" yaml:"windowBefore"`
}

