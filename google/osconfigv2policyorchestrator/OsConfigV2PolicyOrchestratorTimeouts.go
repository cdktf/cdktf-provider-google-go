// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestrator


type OsConfigV2PolicyOrchestratorTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/os_config_v2_policy_orchestrator#create OsConfigV2PolicyOrchestrator#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/os_config_v2_policy_orchestrator#delete OsConfigV2PolicyOrchestrator#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/os_config_v2_policy_orchestrator#update OsConfigV2PolicyOrchestrator#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

