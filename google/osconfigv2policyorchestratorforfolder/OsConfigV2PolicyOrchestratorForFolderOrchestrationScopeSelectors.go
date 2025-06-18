// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorforfolder


type OsConfigV2PolicyOrchestratorForFolderOrchestrationScopeSelectors struct {
	// location_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#location_selector OsConfigV2PolicyOrchestratorForFolder#location_selector}
	LocationSelector *OsConfigV2PolicyOrchestratorForFolderOrchestrationScopeSelectorsLocationSelector `field:"optional" json:"locationSelector" yaml:"locationSelector"`
	// resource_hierarchy_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#resource_hierarchy_selector OsConfigV2PolicyOrchestratorForFolder#resource_hierarchy_selector}
	ResourceHierarchySelector *OsConfigV2PolicyOrchestratorForFolderOrchestrationScopeSelectorsResourceHierarchySelector `field:"optional" json:"resourceHierarchySelector" yaml:"resourceHierarchySelector"`
}

