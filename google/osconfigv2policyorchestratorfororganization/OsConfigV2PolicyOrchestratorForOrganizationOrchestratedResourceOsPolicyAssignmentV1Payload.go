// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorfororganization


type OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1Payload struct {
	// instance_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#instance_filter OsConfigV2PolicyOrchestratorForOrganization#instance_filter}
	InstanceFilter *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadInstanceFilter `field:"required" json:"instanceFilter" yaml:"instanceFilter"`
	// os_policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#os_policies OsConfigV2PolicyOrchestratorForOrganization#os_policies}
	OsPolicies interface{} `field:"required" json:"osPolicies" yaml:"osPolicies"`
	// rollout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#rollout OsConfigV2PolicyOrchestratorForOrganization#rollout}
	Rollout *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadRollout `field:"required" json:"rollout" yaml:"rollout"`
	// OS policy assignment description. Length of the description is limited to 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#description OsConfigV2PolicyOrchestratorForOrganization#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The etag for this OS policy assignment. If this is provided on update, it must match the server's etag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#etag OsConfigV2PolicyOrchestratorForOrganization#etag}
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
	// Resource name.
	//
	// Format:
	// 'projects/{project_number}/locations/{location}/osPolicyAssignments/{os_policy_assignment_id}'
	//
	// This field is ignored when you create an OS policy assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#name OsConfigV2PolicyOrchestratorForOrganization#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

