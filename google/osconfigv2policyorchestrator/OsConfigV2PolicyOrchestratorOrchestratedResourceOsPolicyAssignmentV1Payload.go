// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestrator


type OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1Payload struct {
	// instance_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/os_config_v2_policy_orchestrator#instance_filter OsConfigV2PolicyOrchestrator#instance_filter}
	InstanceFilter *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadInstanceFilter `field:"required" json:"instanceFilter" yaml:"instanceFilter"`
	// os_policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/os_config_v2_policy_orchestrator#os_policies OsConfigV2PolicyOrchestrator#os_policies}
	OsPolicies interface{} `field:"required" json:"osPolicies" yaml:"osPolicies"`
	// rollout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/os_config_v2_policy_orchestrator#rollout OsConfigV2PolicyOrchestrator#rollout}
	Rollout *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadRollout `field:"required" json:"rollout" yaml:"rollout"`
	// OS policy assignment description. Length of the description is limited to 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/os_config_v2_policy_orchestrator#description OsConfigV2PolicyOrchestrator#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Resource name.
	//
	// Format:
	// 'projects/{project_number}/locations/{location}/osPolicyAssignments/{os_policy_assignment_id}'
	//
	// This field is ignored when you create an OS policy assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/os_config_v2_policy_orchestrator#name OsConfigV2PolicyOrchestrator#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

