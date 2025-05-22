// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestrator


type OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesRepository struct {
	// apt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/os_config_v2_policy_orchestrator#apt OsConfigV2PolicyOrchestrator#apt}
	Apt *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesRepositoryApt `field:"optional" json:"apt" yaml:"apt"`
	// goo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/os_config_v2_policy_orchestrator#goo OsConfigV2PolicyOrchestrator#goo}
	Goo *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesRepositoryGoo `field:"optional" json:"goo" yaml:"goo"`
	// yum block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/os_config_v2_policy_orchestrator#yum OsConfigV2PolicyOrchestrator#yum}
	Yum *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesRepositoryYum `field:"optional" json:"yum" yaml:"yum"`
	// zypper block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/os_config_v2_policy_orchestrator#zypper OsConfigV2PolicyOrchestrator#zypper}
	Zypper *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesRepositoryZypper `field:"optional" json:"zypper" yaml:"zypper"`
}

