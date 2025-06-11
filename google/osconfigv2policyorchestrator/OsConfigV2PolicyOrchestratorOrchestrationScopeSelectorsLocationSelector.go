// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestrator


type OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsLocationSelector struct {
	// Optional. Names of the locations in scope. Format: 'us-central1-a'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/os_config_v2_policy_orchestrator#included_locations OsConfigV2PolicyOrchestrator#included_locations}
	IncludedLocations *[]*string `field:"optional" json:"includedLocations" yaml:"includedLocations"`
}

