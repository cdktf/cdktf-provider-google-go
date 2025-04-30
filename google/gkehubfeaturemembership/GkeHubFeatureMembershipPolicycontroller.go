// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeaturemembership


type GkeHubFeatureMembershipPolicycontroller struct {
	// policy_controller_hub_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/gke_hub_feature_membership#policy_controller_hub_config GkeHubFeatureMembership#policy_controller_hub_config}
	PolicyControllerHubConfig *GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfig `field:"required" json:"policyControllerHubConfig" yaml:"policyControllerHubConfig"`
	// Optional. Version of Policy Controller to install. Defaults to the latest version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/gke_hub_feature_membership#version GkeHubFeatureMembership#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

