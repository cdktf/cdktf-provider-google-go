// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeaturemembership


type GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring struct {
	// Specifies the list of backends Policy Controller will export to. Specifying an empty value `[]` disables metrics export.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/gke_hub_feature_membership#backends GkeHubFeatureMembership#backends}
	Backends *[]*string `field:"optional" json:"backends" yaml:"backends"`
}

