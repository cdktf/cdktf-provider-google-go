// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorfororganization

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OsConfigV2PolicyOrchestratorForOrganizationConfig struct {
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
	// Required.
	//
	// Action to be done by the orchestrator in
	// 'projects/{project_id}/zones/{zone_id}' locations defined by the
	// 'orchestration_scope'. Allowed values:
	// - 'UPSERT' - Orchestrator will create or update target resources.
	// - 'DELETE' - Orchestrator will delete target resources, if they exist
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/os_config_v2_policy_orchestrator_for_organization#action OsConfigV2PolicyOrchestratorForOrganization#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// orchestrated_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/os_config_v2_policy_orchestrator_for_organization#orchestrated_resource OsConfigV2PolicyOrchestratorForOrganization#orchestrated_resource}
	OrchestratedResource *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResource `field:"required" json:"orchestratedResource" yaml:"orchestratedResource"`
	// Part of 'parent'. Required. The parent resource name in the form of: * 'organizations/{organization_id}/locations/global' * 'folders/{folder_id}/locations/global' * 'projects/{project_id_or_number}/locations/global'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/os_config_v2_policy_orchestrator_for_organization#organization_id OsConfigV2PolicyOrchestratorForOrganization#organization_id}
	OrganizationId *string `field:"required" json:"organizationId" yaml:"organizationId"`
	// Required. The logical identifier of the policy orchestrator, with the following restrictions:.
	//
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the parent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/os_config_v2_policy_orchestrator_for_organization#policy_orchestrator_id OsConfigV2PolicyOrchestratorForOrganization#policy_orchestrator_id}
	PolicyOrchestratorId *string `field:"required" json:"policyOrchestratorId" yaml:"policyOrchestratorId"`
	// Optional. Freeform text describing the purpose of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/os_config_v2_policy_orchestrator_for_organization#description OsConfigV2PolicyOrchestratorForOrganization#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/os_config_v2_policy_orchestrator_for_organization#id OsConfigV2PolicyOrchestratorForOrganization#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional. Labels as key value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/os_config_v2_policy_orchestrator_for_organization#labels OsConfigV2PolicyOrchestratorForOrganization#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// orchestration_scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/os_config_v2_policy_orchestrator_for_organization#orchestration_scope OsConfigV2PolicyOrchestratorForOrganization#orchestration_scope}
	OrchestrationScope *OsConfigV2PolicyOrchestratorForOrganizationOrchestrationScope `field:"optional" json:"orchestrationScope" yaml:"orchestrationScope"`
	// Optional.
	//
	// State of the orchestrator. Can be updated to change orchestrator behaviour.
	// Allowed values:
	// - 'ACTIVE' - orchestrator is actively looking for actions to be taken.
	// - 'STOPPED' - orchestrator won't make any changes.
	//
	// Note: There might be more states added in the future. We use string here
	// instead of an enum, to avoid the need of propagating new states to all the
	// client code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/os_config_v2_policy_orchestrator_for_organization#state OsConfigV2PolicyOrchestratorForOrganization#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/os_config_v2_policy_orchestrator_for_organization#timeouts OsConfigV2PolicyOrchestratorForOrganization#timeouts}
	Timeouts *OsConfigV2PolicyOrchestratorForOrganizationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

