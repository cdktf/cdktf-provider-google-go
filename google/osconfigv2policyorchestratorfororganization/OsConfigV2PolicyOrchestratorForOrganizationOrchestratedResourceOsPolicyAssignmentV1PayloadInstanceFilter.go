// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorfororganization


type OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadInstanceFilter struct {
	// Target all VMs in the project. If true, no other criteria is permitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#all OsConfigV2PolicyOrchestratorForOrganization#all}
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// exclusion_labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#exclusion_labels OsConfigV2PolicyOrchestratorForOrganization#exclusion_labels}
	ExclusionLabels interface{} `field:"optional" json:"exclusionLabels" yaml:"exclusionLabels"`
	// inclusion_labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#inclusion_labels OsConfigV2PolicyOrchestratorForOrganization#inclusion_labels}
	InclusionLabels interface{} `field:"optional" json:"inclusionLabels" yaml:"inclusionLabels"`
	// inventories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#inventories OsConfigV2PolicyOrchestratorForOrganization#inventories}
	Inventories interface{} `field:"optional" json:"inventories" yaml:"inventories"`
}

