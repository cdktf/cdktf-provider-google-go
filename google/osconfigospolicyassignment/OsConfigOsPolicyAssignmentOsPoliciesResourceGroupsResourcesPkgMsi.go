// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigospolicyassignment


type OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi struct {
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/os_config_os_policy_assignment#source OsConfigOsPolicyAssignment#source}
	Source *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsiSource `field:"required" json:"source" yaml:"source"`
	// Additional properties to use during installation.
	//
	// This should be in the format of Property=Setting. Appended to the defaults of 'ACTION=INSTALL REBOOT=ReallySuppress'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/os_config_os_policy_assignment#properties OsConfigOsPolicyAssignment#properties}
	Properties *[]*string `field:"optional" json:"properties" yaml:"properties"`
}

