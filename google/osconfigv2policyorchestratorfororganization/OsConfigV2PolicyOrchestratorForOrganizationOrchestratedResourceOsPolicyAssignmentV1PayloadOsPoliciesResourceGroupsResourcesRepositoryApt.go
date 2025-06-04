// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorfororganization


type OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesRepositoryApt struct {
	// Required. Type of archive files in this repository. Possible values: ARCHIVE_TYPE_UNSPECIFIED DEB DEB_SRC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#archive_type OsConfigV2PolicyOrchestratorForOrganization#archive_type}
	ArchiveType *string `field:"required" json:"archiveType" yaml:"archiveType"`
	// Required. List of components for this repository. Must contain at least one item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#components OsConfigV2PolicyOrchestratorForOrganization#components}
	Components *[]*string `field:"required" json:"components" yaml:"components"`
	// Required. Distribution of this repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#distribution OsConfigV2PolicyOrchestratorForOrganization#distribution}
	Distribution *string `field:"required" json:"distribution" yaml:"distribution"`
	// Required. URI for this repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#uri OsConfigV2PolicyOrchestratorForOrganization#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// URI of the key file for this repository. The agent maintains a keyring at '/etc/apt/trusted.gpg.d/osconfig_agent_managed.gpg'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#gpg_key OsConfigV2PolicyOrchestratorForOrganization#gpg_key}
	GpgKey *string `field:"optional" json:"gpgKey" yaml:"gpgKey"`
}

