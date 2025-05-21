// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorforfolder

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OsConfigV2PolicyOrchestratorForFolderConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Action to be done by the orchestrator in 'projects/{project_id}/zones/{zone_id}' locations defined by the 'orchestration_scope'.
	//
	// Allowed values:
	// - 'UPSERT' - Orchestrator will create or update target resources.
	// - 'DELETE' - Orchestrator will delete target resources, if they exist
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#action OsConfigV2PolicyOrchestratorForFolder#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// The parent resource name in the form of 'folders/{folder_id}/locations/global'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#folder_id OsConfigV2PolicyOrchestratorForFolder#folder_id}
	FolderId *string `field:"required" json:"folderId" yaml:"folderId"`
	// orchestrated_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#orchestrated_resource OsConfigV2PolicyOrchestratorForFolder#orchestrated_resource}
	OrchestratedResource *OsConfigV2PolicyOrchestratorForFolderOrchestratedResource `field:"required" json:"orchestratedResource" yaml:"orchestratedResource"`
	// The logical identifier of the policy orchestrator, with the following restrictions:.
	//
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the parent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#policy_orchestrator_id OsConfigV2PolicyOrchestratorForFolder#policy_orchestrator_id}
	PolicyOrchestratorId *string `field:"required" json:"policyOrchestratorId" yaml:"policyOrchestratorId"`
	// Freeform text describing the purpose of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#description OsConfigV2PolicyOrchestratorForFolder#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#id OsConfigV2PolicyOrchestratorForFolder#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#labels OsConfigV2PolicyOrchestratorForFolder#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// orchestration_scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#orchestration_scope OsConfigV2PolicyOrchestratorForFolder#orchestration_scope}
	OrchestrationScope *OsConfigV2PolicyOrchestratorForFolderOrchestrationScope `field:"optional" json:"orchestrationScope" yaml:"orchestrationScope"`
	// State of the orchestrator.
	//
	// Can be updated to change orchestrator behaviour.
	// Allowed values:
	// - 'ACTIVE' - orchestrator is actively looking for actions to be taken.
	// - 'STOPPED' - orchestrator won't make any changes.
	//
	// Note: There might be more states added in the future. We use string here
	// instead of an enum, to avoid the need of propagating new states to all the
	// client code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#state OsConfigV2PolicyOrchestratorForFolder#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#timeouts OsConfigV2PolicyOrchestratorForFolder#timeouts}
	Timeouts *OsConfigV2PolicyOrchestratorForFolderTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

