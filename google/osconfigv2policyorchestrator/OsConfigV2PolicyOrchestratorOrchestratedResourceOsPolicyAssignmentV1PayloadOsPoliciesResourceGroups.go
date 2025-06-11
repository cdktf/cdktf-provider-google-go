// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestrator


type OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroups struct {
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/os_config_v2_policy_orchestrator#resources OsConfigV2PolicyOrchestrator#resources}
	Resources interface{} `field:"required" json:"resources" yaml:"resources"`
	// inventory_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/os_config_v2_policy_orchestrator#inventory_filters OsConfigV2PolicyOrchestrator#inventory_filters}
	InventoryFilters interface{} `field:"optional" json:"inventoryFilters" yaml:"inventoryFilters"`
}

