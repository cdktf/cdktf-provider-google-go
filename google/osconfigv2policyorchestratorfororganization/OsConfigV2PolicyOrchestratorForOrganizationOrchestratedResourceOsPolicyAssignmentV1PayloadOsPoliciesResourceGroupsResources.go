// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorfororganization


type OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResources struct {
	// Required. The id of the resource with the following restrictions:.
	//
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the OS policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#id OsConfigV2PolicyOrchestratorForOrganization#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// exec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#exec OsConfigV2PolicyOrchestratorForOrganization#exec}
	Exec *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesExec `field:"optional" json:"exec" yaml:"exec"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#file OsConfigV2PolicyOrchestratorForOrganization#file}
	File *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesFile `field:"optional" json:"file" yaml:"file"`
	// pkg block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#pkg OsConfigV2PolicyOrchestratorForOrganization#pkg}
	Pkg *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkg `field:"optional" json:"pkg" yaml:"pkg"`
	// repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#repository OsConfigV2PolicyOrchestratorForOrganization#repository}
	Repository *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesRepository `field:"optional" json:"repository" yaml:"repository"`
}

