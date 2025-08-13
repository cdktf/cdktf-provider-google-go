// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestrator


type OsConfigV2PolicyOrchestratorOrchestrationScopeSelectors struct {
	// location_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/os_config_v2_policy_orchestrator#location_selector OsConfigV2PolicyOrchestrator#location_selector}
	LocationSelector *OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsLocationSelector `field:"optional" json:"locationSelector" yaml:"locationSelector"`
	// resource_hierarchy_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/os_config_v2_policy_orchestrator#resource_hierarchy_selector OsConfigV2PolicyOrchestrator#resource_hierarchy_selector}
	ResourceHierarchySelector *OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsResourceHierarchySelector `field:"optional" json:"resourceHierarchySelector" yaml:"resourceHierarchySelector"`
}

