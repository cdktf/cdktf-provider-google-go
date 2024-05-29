// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeature


type GkeHubFeatureSpec struct {
	// clusterupgrade block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/gke_hub_feature#clusterupgrade GkeHubFeature#clusterupgrade}
	Clusterupgrade *GkeHubFeatureSpecClusterupgrade `field:"optional" json:"clusterupgrade" yaml:"clusterupgrade"`
	// fleetobservability block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/gke_hub_feature#fleetobservability GkeHubFeature#fleetobservability}
	Fleetobservability *GkeHubFeatureSpecFleetobservability `field:"optional" json:"fleetobservability" yaml:"fleetobservability"`
	// multiclusteringress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/gke_hub_feature#multiclusteringress GkeHubFeature#multiclusteringress}
	Multiclusteringress *GkeHubFeatureSpecMulticlusteringress `field:"optional" json:"multiclusteringress" yaml:"multiclusteringress"`
}

