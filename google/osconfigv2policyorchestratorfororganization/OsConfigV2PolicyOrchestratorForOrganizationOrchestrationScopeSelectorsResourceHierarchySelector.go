// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorfororganization


type OsConfigV2PolicyOrchestratorForOrganizationOrchestrationScopeSelectorsResourceHierarchySelector struct {
	// Optional. Names of the folders in scope. Format: 'folders/{folder_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/os_config_v2_policy_orchestrator_for_organization#included_folders OsConfigV2PolicyOrchestratorForOrganization#included_folders}
	IncludedFolders *[]*string `field:"optional" json:"includedFolders" yaml:"includedFolders"`
	// Optional. Names of the projects in scope. Format: 'projects/{project_number}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/os_config_v2_policy_orchestrator_for_organization#included_projects OsConfigV2PolicyOrchestratorForOrganization#included_projects}
	IncludedProjects *[]*string `field:"optional" json:"includedProjects" yaml:"includedProjects"`
}

