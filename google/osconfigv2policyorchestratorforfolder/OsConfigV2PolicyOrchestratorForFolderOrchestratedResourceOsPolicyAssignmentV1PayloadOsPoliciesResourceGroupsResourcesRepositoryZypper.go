// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorforfolder


type OsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesRepositoryZypper struct {
	// The location of the repository directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#base_url OsConfigV2PolicyOrchestratorForFolder#base_url}
	BaseUrl *string `field:"required" json:"baseUrl" yaml:"baseUrl"`
	// A one word, unique name for this repository.
	//
	// This is the 'repo
	// id' in the zypper config file and also the 'display_name' if
	// 'display_name' is omitted. This id is also used as the unique
	// identifier when checking for GuestPolicy conflicts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#id OsConfigV2PolicyOrchestratorForFolder#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The display name of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#display_name OsConfigV2PolicyOrchestratorForFolder#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// URIs of GPG keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#gpg_keys OsConfigV2PolicyOrchestratorForFolder#gpg_keys}
	GpgKeys *[]*string `field:"optional" json:"gpgKeys" yaml:"gpgKeys"`
}

