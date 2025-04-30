// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorfororganization


type OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesRepositoryGoo struct {
	// Required. The name of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#name OsConfigV2PolicyOrchestratorForOrganization#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Required. The url of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#url OsConfigV2PolicyOrchestratorForOrganization#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

