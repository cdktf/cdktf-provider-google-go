// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeaturemembership


type GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodTolerations struct {
	// Matches a taint effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/gke_hub_feature_membership#effect GkeHubFeatureMembership#effect}
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// Matches a taint key (not necessarily unique).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/gke_hub_feature_membership#key GkeHubFeatureMembership#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Matches a taint operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/gke_hub_feature_membership#operator GkeHubFeatureMembership#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Matches a taint value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/gke_hub_feature_membership#value GkeHubFeatureMembership#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

