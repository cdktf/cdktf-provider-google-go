// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorfororganization


type OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadRolloutDisruptionBudget struct {
	// Specifies a fixed value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#fixed OsConfigV2PolicyOrchestratorForOrganization#fixed}
	Fixed *float64 `field:"optional" json:"fixed" yaml:"fixed"`
	// Specifies the relative value defined as a percentage, which will be multiplied by a reference value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#percent OsConfigV2PolicyOrchestratorForOrganization#percent}
	Percent *float64 `field:"optional" json:"percent" yaml:"percent"`
}

