// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorforfolder


type OsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesRepositoryApt struct {
	// Type of archive files in this repository. Possible values: ["DEB", "DEB_SRC"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#archive_type OsConfigV2PolicyOrchestratorForFolder#archive_type}
	ArchiveType *string `field:"required" json:"archiveType" yaml:"archiveType"`
	// List of components for this repository. Must contain at least one item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#components OsConfigV2PolicyOrchestratorForFolder#components}
	Components *[]*string `field:"required" json:"components" yaml:"components"`
	// Distribution of this repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#distribution OsConfigV2PolicyOrchestratorForFolder#distribution}
	Distribution *string `field:"required" json:"distribution" yaml:"distribution"`
	// URI for this repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#uri OsConfigV2PolicyOrchestratorForFolder#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// URI of the key file for this repository. The agent maintains a keyring at '/etc/apt/trusted.gpg.d/osconfig_agent_managed.gpg'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#gpg_key OsConfigV2PolicyOrchestratorForFolder#gpg_key}
	GpgKey *string `field:"optional" json:"gpgKey" yaml:"gpgKey"`
}

