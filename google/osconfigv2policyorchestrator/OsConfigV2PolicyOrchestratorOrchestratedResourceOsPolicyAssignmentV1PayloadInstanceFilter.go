// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestrator


type OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadInstanceFilter struct {
	// Target all VMs in the project. If true, no other criteria is permitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/os_config_v2_policy_orchestrator#all OsConfigV2PolicyOrchestrator#all}
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// exclusion_labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/os_config_v2_policy_orchestrator#exclusion_labels OsConfigV2PolicyOrchestrator#exclusion_labels}
	ExclusionLabels interface{} `field:"optional" json:"exclusionLabels" yaml:"exclusionLabels"`
	// inclusion_labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/os_config_v2_policy_orchestrator#inclusion_labels OsConfigV2PolicyOrchestrator#inclusion_labels}
	InclusionLabels interface{} `field:"optional" json:"inclusionLabels" yaml:"inclusionLabels"`
	// inventories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/os_config_v2_policy_orchestrator#inventories OsConfigV2PolicyOrchestrator#inventories}
	Inventories interface{} `field:"optional" json:"inventories" yaml:"inventories"`
}

