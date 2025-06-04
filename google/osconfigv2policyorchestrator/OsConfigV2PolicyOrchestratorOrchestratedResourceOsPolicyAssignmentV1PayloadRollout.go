// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestrator


type OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadRollout struct {
	// disruption_budget block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/os_config_v2_policy_orchestrator#disruption_budget OsConfigV2PolicyOrchestrator#disruption_budget}
	DisruptionBudget *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadRolloutDisruptionBudget `field:"required" json:"disruptionBudget" yaml:"disruptionBudget"`
	// Required.
	//
	// This determines the minimum duration of time to wait after the
	// configuration changes are applied through the current rollout. A
	// VM continues to count towards the 'disruption_budget' at least
	// until this duration of time has passed after configuration changes are
	// applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/os_config_v2_policy_orchestrator#min_wait_duration OsConfigV2PolicyOrchestrator#min_wait_duration}
	MinWaitDuration *string `field:"required" json:"minWaitDuration" yaml:"minWaitDuration"`
}

