// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorforfolder


type OsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroups struct {
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#resources OsConfigV2PolicyOrchestratorForFolder#resources}
	Resources interface{} `field:"required" json:"resources" yaml:"resources"`
	// inventory_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#inventory_filters OsConfigV2PolicyOrchestratorForFolder#inventory_filters}
	InventoryFilters interface{} `field:"optional" json:"inventoryFilters" yaml:"inventoryFilters"`
}

