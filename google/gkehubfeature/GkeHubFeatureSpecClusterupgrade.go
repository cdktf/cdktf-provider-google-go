// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeature


type GkeHubFeatureSpecClusterupgrade struct {
	// Specified if other fleet should be considered as a source of upgrades.
	//
	// Currently, at most one upstream fleet is allowed. The fleet name should be either fleet project number or id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/gke_hub_feature#upstream_fleets GkeHubFeature#upstream_fleets}
	UpstreamFleets *[]*string `field:"required" json:"upstreamFleets" yaml:"upstreamFleets"`
	// gke_upgrade_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/gke_hub_feature#gke_upgrade_overrides GkeHubFeature#gke_upgrade_overrides}
	GkeUpgradeOverrides interface{} `field:"optional" json:"gkeUpgradeOverrides" yaml:"gkeUpgradeOverrides"`
	// post_conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/gke_hub_feature#post_conditions GkeHubFeature#post_conditions}
	PostConditions *GkeHubFeatureSpecClusterupgradePostConditions `field:"optional" json:"postConditions" yaml:"postConditions"`
}

