// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigospolicyassignment


type OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo struct {
	// The name of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/os_config_os_policy_assignment#name OsConfigOsPolicyAssignment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The url of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/os_config_os_policy_assignment#url OsConfigOsPolicyAssignment#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

