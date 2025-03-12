// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigospolicyassignment


type OsConfigOsPolicyAssignmentOsPoliciesResourceGroups struct {
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/os_config_os_policy_assignment#resources OsConfigOsPolicyAssignment#resources}
	Resources interface{} `field:"required" json:"resources" yaml:"resources"`
	// inventory_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/os_config_os_policy_assignment#inventory_filters OsConfigOsPolicyAssignment#inventory_filters}
	InventoryFilters interface{} `field:"optional" json:"inventoryFilters" yaml:"inventoryFilters"`
}

