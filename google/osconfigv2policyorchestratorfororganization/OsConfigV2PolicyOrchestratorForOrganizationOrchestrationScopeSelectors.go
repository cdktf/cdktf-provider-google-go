// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorfororganization


type OsConfigV2PolicyOrchestratorForOrganizationOrchestrationScopeSelectors struct {
	// location_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#location_selector OsConfigV2PolicyOrchestratorForOrganization#location_selector}
	LocationSelector *OsConfigV2PolicyOrchestratorForOrganizationOrchestrationScopeSelectorsLocationSelector `field:"optional" json:"locationSelector" yaml:"locationSelector"`
	// resource_hierarchy_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#resource_hierarchy_selector OsConfigV2PolicyOrchestratorForOrganization#resource_hierarchy_selector}
	ResourceHierarchySelector *OsConfigV2PolicyOrchestratorForOrganizationOrchestrationScopeSelectorsResourceHierarchySelector `field:"optional" json:"resourceHierarchySelector" yaml:"resourceHierarchySelector"`
}

