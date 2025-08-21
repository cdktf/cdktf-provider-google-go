// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorfororganization


type OsConfigV2PolicyOrchestratorForOrganizationOrchestrationScopeSelectorsLocationSelector struct {
	// Optional. Names of the locations in scope. Format: 'us-central1-a'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/os_config_v2_policy_orchestrator_for_organization#included_locations OsConfigV2PolicyOrchestratorForOrganization#included_locations}
	IncludedLocations *[]*string `field:"optional" json:"includedLocations" yaml:"includedLocations"`
}

