// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestrator


type OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesFile struct {
	// Required. The absolute path of the file within the VM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/os_config_v2_policy_orchestrator#path OsConfigV2PolicyOrchestrator#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Required. Desired state of the file. Possible values: DESIRED_STATE_UNSPECIFIED PRESENT ABSENT CONTENTS_MATCH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/os_config_v2_policy_orchestrator#state OsConfigV2PolicyOrchestrator#state}
	State *string `field:"required" json:"state" yaml:"state"`
	// A a file with this content. The size of the content is limited to 32KiB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/os_config_v2_policy_orchestrator#content OsConfigV2PolicyOrchestrator#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/os_config_v2_policy_orchestrator#file OsConfigV2PolicyOrchestrator#file}
	File *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesFileFile `field:"optional" json:"file" yaml:"file"`
	// Consists of three octal digits which represent, in order, the permissions of the owner, group, and other users for the file (similarly to the numeric mode used in the linux chmod utility).
	//
	// Each digit represents a three bit number with the 4 bit
	// corresponding to the read permissions, the 2 bit corresponds to the
	// write bit, and the one bit corresponds to the execute permission.
	// Default behavior is 755.
	//
	// Below are some examples of permissions and their associated values:
	// read, write, and execute: 7
	// read and execute: 5
	// read and write: 6
	// read only: 4
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/os_config_v2_policy_orchestrator#permissions OsConfigV2PolicyOrchestrator#permissions}
	Permissions *string `field:"optional" json:"permissions" yaml:"permissions"`
}

