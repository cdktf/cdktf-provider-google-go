// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeature


type GkeHubFeatureFleetDefaultMemberConfig struct {
	// configmanagement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.7.0/docs/resources/gke_hub_feature#configmanagement GkeHubFeature#configmanagement}
	Configmanagement *GkeHubFeatureFleetDefaultMemberConfigConfigmanagement `field:"optional" json:"configmanagement" yaml:"configmanagement"`
	// mesh block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.7.0/docs/resources/gke_hub_feature#mesh GkeHubFeature#mesh}
	Mesh *GkeHubFeatureFleetDefaultMemberConfigMesh `field:"optional" json:"mesh" yaml:"mesh"`
}

