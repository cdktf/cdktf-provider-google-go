// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestrator


type OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesRepositoryApt struct {
	// Required. Type of archive files in this repository. Possible values: ARCHIVE_TYPE_UNSPECIFIED DEB DEB_SRC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/os_config_v2_policy_orchestrator#archive_type OsConfigV2PolicyOrchestrator#archive_type}
	ArchiveType *string `field:"required" json:"archiveType" yaml:"archiveType"`
	// Required. List of components for this repository. Must contain at least one item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/os_config_v2_policy_orchestrator#components OsConfigV2PolicyOrchestrator#components}
	Components *[]*string `field:"required" json:"components" yaml:"components"`
	// Required. Distribution of this repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/os_config_v2_policy_orchestrator#distribution OsConfigV2PolicyOrchestrator#distribution}
	Distribution *string `field:"required" json:"distribution" yaml:"distribution"`
	// Required. URI for this repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/os_config_v2_policy_orchestrator#uri OsConfigV2PolicyOrchestrator#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// URI of the key file for this repository. The agent maintains a keyring at '/etc/apt/trusted.gpg.d/osconfig_agent_managed.gpg'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/os_config_v2_policy_orchestrator#gpg_key OsConfigV2PolicyOrchestrator#gpg_key}
	GpgKey *string `field:"optional" json:"gpgKey" yaml:"gpgKey"`
}

