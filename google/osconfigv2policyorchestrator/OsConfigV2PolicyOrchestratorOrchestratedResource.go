// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestrator


type OsConfigV2PolicyOrchestratorOrchestratedResource struct {
	// Optional. ID of the resource to be used while generating set of affected resources.
	//
	// For UPSERT action the value is auto-generated during PolicyOrchestrator
	// creation when not set. When the value is set it should following next
	// restrictions:
	//
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the project.
	//
	// For DELETE action, ID must be specified explicitly during
	// PolicyOrchestrator creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/os_config_v2_policy_orchestrator#id OsConfigV2PolicyOrchestrator#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// os_policy_assignment_v1_payload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/os_config_v2_policy_orchestrator#os_policy_assignment_v1_payload OsConfigV2PolicyOrchestrator#os_policy_assignment_v1_payload}
	OsPolicyAssignmentV1Payload *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1Payload `field:"optional" json:"osPolicyAssignmentV1Payload" yaml:"osPolicyAssignmentV1Payload"`
}

