// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfleet


type GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfig struct {
	// Mode of operation for binauthz policy evaluation. Possible values: ["DISABLED", "POLICY_BINDINGS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/gke_hub_fleet#evaluation_mode GkeHubFleet#evaluation_mode}
	EvaluationMode *string `field:"optional" json:"evaluationMode" yaml:"evaluationMode"`
	// policy_bindings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/gke_hub_fleet#policy_bindings GkeHubFleet#policy_bindings}
	PolicyBindings interface{} `field:"optional" json:"policyBindings" yaml:"policyBindings"`
}

