// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorfororganization


type OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadInstanceFilterInventories struct {
	// Required. The OS short name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#os_short_name OsConfigV2PolicyOrchestratorForOrganization#os_short_name}
	OsShortName *string `field:"required" json:"osShortName" yaml:"osShortName"`
	// The OS version.
	//
	// Prefix matches are supported if asterisk(*) is provided as the
	// last character. For example, to match all versions with a major
	// version of '7', specify the following value for this field '7.*'
	//
	// An empty string matches all OS versions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#os_version OsConfigV2PolicyOrchestratorForOrganization#os_version}
	OsVersion *string `field:"optional" json:"osVersion" yaml:"osVersion"`
}

