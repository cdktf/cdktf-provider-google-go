// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestrator


type OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesRepositoryGoo struct {
	// Required. The name of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/os_config_v2_policy_orchestrator#name OsConfigV2PolicyOrchestrator#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Required. The url of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/os_config_v2_policy_orchestrator#url OsConfigV2PolicyOrchestrator#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

