// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeature


type GkeHubFeatureFleetDefaultMemberConfigPolicycontroller struct {
	// policy_controller_hub_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/gke_hub_feature#policy_controller_hub_config GkeHubFeature#policy_controller_hub_config}
	PolicyControllerHubConfig *GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfig `field:"required" json:"policyControllerHubConfig" yaml:"policyControllerHubConfig"`
	// Configures the version of Policy Controller.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/gke_hub_feature#version GkeHubFeature#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

