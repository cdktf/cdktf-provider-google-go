// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorforfolder


type OsConfigV2PolicyOrchestratorForFolderOrchestrationScopeSelectorsResourceHierarchySelector struct {
	// Names of the folders in scope. Format: 'folders/{folder_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#included_folders OsConfigV2PolicyOrchestratorForFolder#included_folders}
	IncludedFolders *[]*string `field:"optional" json:"includedFolders" yaml:"includedFolders"`
	// Names of the projects in scope. Format: 'projects/{project_number}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#included_projects OsConfigV2PolicyOrchestratorForFolder#included_projects}
	IncludedProjects *[]*string `field:"optional" json:"includedProjects" yaml:"includedProjects"`
}

