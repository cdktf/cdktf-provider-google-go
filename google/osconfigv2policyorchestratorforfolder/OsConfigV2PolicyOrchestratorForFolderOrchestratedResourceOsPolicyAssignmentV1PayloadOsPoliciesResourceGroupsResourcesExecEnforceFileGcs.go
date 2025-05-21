// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorforfolder


type OsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesExecEnforceFileGcs struct {
	// Bucket of the Cloud Storage object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#bucket OsConfigV2PolicyOrchestratorForFolder#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Name of the Cloud Storage object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#object OsConfigV2PolicyOrchestratorForFolder#object}
	Object *string `field:"required" json:"object" yaml:"object"`
	// Generation number of the Cloud Storage object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#generation OsConfigV2PolicyOrchestratorForFolder#generation}
	Generation *string `field:"optional" json:"generation" yaml:"generation"`
}

