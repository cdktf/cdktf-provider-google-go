// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestrator


type OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkg struct {
	// Required. The desired state the agent should maintain for this package. Possible values: DESIRED_STATE_UNSPECIFIED INSTALLED REMOVED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/os_config_v2_policy_orchestrator#desired_state OsConfigV2PolicyOrchestrator#desired_state}
	DesiredState *string `field:"required" json:"desiredState" yaml:"desiredState"`
	// apt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/os_config_v2_policy_orchestrator#apt OsConfigV2PolicyOrchestrator#apt}
	Apt *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgApt `field:"optional" json:"apt" yaml:"apt"`
	// deb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/os_config_v2_policy_orchestrator#deb OsConfigV2PolicyOrchestrator#deb}
	Deb *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDeb `field:"optional" json:"deb" yaml:"deb"`
	// googet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/os_config_v2_policy_orchestrator#googet OsConfigV2PolicyOrchestrator#googet}
	Googet *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgGooget `field:"optional" json:"googet" yaml:"googet"`
	// msi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/os_config_v2_policy_orchestrator#msi OsConfigV2PolicyOrchestrator#msi}
	Msi *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgMsi `field:"optional" json:"msi" yaml:"msi"`
	// rpm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/os_config_v2_policy_orchestrator#rpm OsConfigV2PolicyOrchestrator#rpm}
	Rpm *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgRpm `field:"optional" json:"rpm" yaml:"rpm"`
	// yum block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/os_config_v2_policy_orchestrator#yum OsConfigV2PolicyOrchestrator#yum}
	Yum *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgYum `field:"optional" json:"yum" yaml:"yum"`
	// zypper block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/os_config_v2_policy_orchestrator#zypper OsConfigV2PolicyOrchestrator#zypper}
	Zypper *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgZypper `field:"optional" json:"zypper" yaml:"zypper"`
}

