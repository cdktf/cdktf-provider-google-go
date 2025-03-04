// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeature


type GkeHubFeatureSpecClusterupgradeGkeUpgradeOverridesUpgrade struct {
	// Name of the upgrade, e.g., "k8s_control_plane". It should be a valid upgrade name. It must not exceet 99 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/gke_hub_feature#name GkeHubFeature#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Version of the upgrade, e.g., "1.22.1-gke.100". It should be a valid version. It must not exceet 99 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/gke_hub_feature#version GkeHubFeature#version}
	Version *string `field:"required" json:"version" yaml:"version"`
}

