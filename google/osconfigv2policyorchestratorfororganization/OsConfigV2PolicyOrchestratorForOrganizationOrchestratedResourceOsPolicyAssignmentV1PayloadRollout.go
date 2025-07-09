// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorfororganization


type OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadRollout struct {
	// disruption_budget block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#disruption_budget OsConfigV2PolicyOrchestratorForOrganization#disruption_budget}
	DisruptionBudget *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadRolloutDisruptionBudget `field:"required" json:"disruptionBudget" yaml:"disruptionBudget"`
	// Required.
	//
	// This determines the minimum duration of time to wait after the
	// configuration changes are applied through the current rollout. A
	// VM continues to count towards the 'disruption_budget' at least
	// until this duration of time has passed after configuration changes are
	// applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#min_wait_duration OsConfigV2PolicyOrchestratorForOrganization#min_wait_duration}
	MinWaitDuration *string `field:"required" json:"minWaitDuration" yaml:"minWaitDuration"`
}

