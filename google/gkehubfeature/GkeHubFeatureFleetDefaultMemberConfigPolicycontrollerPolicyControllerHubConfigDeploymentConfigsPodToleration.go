// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeature


type GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodToleration struct {
	// Matches a taint effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.12.0/docs/resources/gke_hub_feature#effect GkeHubFeature#effect}
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// Matches a taint key (not necessarily unique).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.12.0/docs/resources/gke_hub_feature#key GkeHubFeature#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Matches a taint operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.12.0/docs/resources/gke_hub_feature#operator GkeHubFeature#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Matches a taint value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.12.0/docs/resources/gke_hub_feature#value GkeHubFeature#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

