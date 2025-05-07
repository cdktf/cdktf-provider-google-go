// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorforfolder


type OsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1Payload struct {
	// instance_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#instance_filter OsConfigV2PolicyOrchestratorForFolder#instance_filter}
	InstanceFilter *OsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadInstanceFilter `field:"required" json:"instanceFilter" yaml:"instanceFilter"`
	// os_policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#os_policies OsConfigV2PolicyOrchestratorForFolder#os_policies}
	OsPolicies interface{} `field:"required" json:"osPolicies" yaml:"osPolicies"`
	// rollout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#rollout OsConfigV2PolicyOrchestratorForFolder#rollout}
	Rollout *OsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadRollout `field:"required" json:"rollout" yaml:"rollout"`
	// OS policy assignment description. Length of the description is limited to 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#description OsConfigV2PolicyOrchestratorForFolder#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Resource name.
	//
	// Format:
	// 'projects/{project_number}/locations/{location}/osPolicyAssignments/{os_policy_assignment_id}'
	//
	// This field is ignored when you create an OS policy assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#name OsConfigV2PolicyOrchestratorForFolder#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

