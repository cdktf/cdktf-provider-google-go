// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeature


type GkeHubFeatureFleetDefaultMemberConfigConfigmanagement struct {
	// config_sync block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/gke_hub_feature#config_sync GkeHubFeature#config_sync}
	ConfigSync *GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSync `field:"optional" json:"configSync" yaml:"configSync"`
	// Version of ACM installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/gke_hub_feature#version GkeHubFeature#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

