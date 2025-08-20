// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorforfolder


type OsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgRpmSource struct {
	// Defaults to false. When false, files are subject to validations based on the file type:.
	//
	// Remote: A checksum must be specified.
	// Cloud Storage: An object generation number must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#allow_insecure OsConfigV2PolicyOrchestratorForFolder#allow_insecure}
	AllowInsecure interface{} `field:"optional" json:"allowInsecure" yaml:"allowInsecure"`
	// gcs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#gcs OsConfigV2PolicyOrchestratorForFolder#gcs}
	Gcs *OsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgRpmSourceGcs `field:"optional" json:"gcs" yaml:"gcs"`
	// A local path within the VM to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#local_path OsConfigV2PolicyOrchestratorForFolder#local_path}
	LocalPath *string `field:"optional" json:"localPath" yaml:"localPath"`
	// remote block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#remote OsConfigV2PolicyOrchestratorForFolder#remote}
	Remote *OsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgRpmSourceRemote `field:"optional" json:"remote" yaml:"remote"`
}

