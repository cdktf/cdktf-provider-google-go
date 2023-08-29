// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigospolicyassignment


type OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec struct {
	// validate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.80.0/docs/resources/os_config_os_policy_assignment#validate OsConfigOsPolicyAssignment#validate}
	Validate *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecValidate `field:"required" json:"validate" yaml:"validate"`
	// enforce block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.80.0/docs/resources/os_config_os_policy_assignment#enforce OsConfigOsPolicyAssignment#enforce}
	Enforce *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecEnforce `field:"optional" json:"enforce" yaml:"enforce"`
}

