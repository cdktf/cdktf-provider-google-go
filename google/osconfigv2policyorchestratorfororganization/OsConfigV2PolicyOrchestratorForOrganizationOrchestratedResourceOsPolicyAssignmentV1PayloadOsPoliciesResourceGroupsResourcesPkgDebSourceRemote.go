// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorfororganization


type OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceRemote struct {
	// Required. URI from which to fetch the object. It should contain both the protocol and path following the format '{protocol}://{location}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#uri OsConfigV2PolicyOrchestratorForOrganization#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// SHA256 checksum of the remote file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#sha256_checksum OsConfigV2PolicyOrchestratorForOrganization#sha256_checksum}
	Sha256Checksum *string `field:"optional" json:"sha256Checksum" yaml:"sha256Checksum"`
}

