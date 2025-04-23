// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestrator


type OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesExec struct {
	// validate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/os_config_v2_policy_orchestrator#validate OsConfigV2PolicyOrchestrator#validate}
	Validate *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesExecValidate `field:"required" json:"validate" yaml:"validate"`
	// enforce block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/os_config_v2_policy_orchestrator#enforce OsConfigV2PolicyOrchestrator#enforce}
	Enforce *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesExecEnforce `field:"optional" json:"enforce" yaml:"enforce"`
}

