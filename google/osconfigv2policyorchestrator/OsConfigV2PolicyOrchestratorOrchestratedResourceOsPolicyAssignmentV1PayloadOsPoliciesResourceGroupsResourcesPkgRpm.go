// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestrator


type OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgRpm struct {
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/os_config_v2_policy_orchestrator#source OsConfigV2PolicyOrchestrator#source}
	Source *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgRpmSource `field:"required" json:"source" yaml:"source"`
	// Whether dependencies should also be installed.
	//
	// - install when false: 'rpm --upgrade --replacepkgs package.rpm'
	// - install when true: 'yum -y install package.rpm' or
	// 'zypper -y install package.rpm'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/os_config_v2_policy_orchestrator#pull_deps OsConfigV2PolicyOrchestrator#pull_deps}
	PullDeps interface{} `field:"optional" json:"pullDeps" yaml:"pullDeps"`
}

