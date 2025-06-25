// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorforfolder


type OsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesExecEnforce struct {
	// The script interpreter to use. Possible values: ["NONE", "SHELL", "POWERSHELL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#interpreter OsConfigV2PolicyOrchestratorForFolder#interpreter}
	Interpreter *string `field:"required" json:"interpreter" yaml:"interpreter"`
	// Optional arguments to pass to the source during execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#args OsConfigV2PolicyOrchestratorForFolder#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#file OsConfigV2PolicyOrchestratorForFolder#file}
	File *OsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesExecEnforceFile `field:"optional" json:"file" yaml:"file"`
	// Only recorded for enforce Exec.
	//
	// Path to an output file (that is created by this Exec) whose
	// content will be recorded in OSPolicyResourceCompliance after a
	// successful run. Absence or failure to read this file will result in
	// this ExecResource being non-compliant. Output file size is limited to
	// 500K bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#output_file_path OsConfigV2PolicyOrchestratorForFolder#output_file_path}
	OutputFilePath *string `field:"optional" json:"outputFilePath" yaml:"outputFilePath"`
	// An inline script. The size of the script is limited to 32KiB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#script OsConfigV2PolicyOrchestratorForFolder#script}
	Script *string `field:"optional" json:"script" yaml:"script"`
}

