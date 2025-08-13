// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorfororganization


type OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgMsiSourceGcs struct {
	// Required. Bucket of the Cloud Storage object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#bucket OsConfigV2PolicyOrchestratorForOrganization#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Required. Name of the Cloud Storage object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#object OsConfigV2PolicyOrchestratorForOrganization#object}
	Object *string `field:"required" json:"object" yaml:"object"`
	// Generation number of the Cloud Storage object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#generation OsConfigV2PolicyOrchestratorForOrganization#generation}
	Generation *string `field:"optional" json:"generation" yaml:"generation"`
}

