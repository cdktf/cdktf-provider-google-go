// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeature


type GkeHubFeatureFleetDefaultMemberConfigConfigmanagement struct {
	// config_sync block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/gke_hub_feature#config_sync GkeHubFeature#config_sync}
	ConfigSync *GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSync `field:"optional" json:"configSync" yaml:"configSync"`
	// Set this field to MANAGEMENT_AUTOMATIC to enable Config Sync auto-upgrades, and set this field to MANAGEMENT_MANUAL or MANAGEMENT_UNSPECIFIED to disable Config Sync auto-upgrades.
	//
	// Possible values: ["MANAGEMENT_UNSPECIFIED", "MANAGEMENT_AUTOMATIC", "MANAGEMENT_MANUAL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/gke_hub_feature#management GkeHubFeature#management}
	Management *string `field:"optional" json:"management" yaml:"management"`
	// Version of Config Sync installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/gke_hub_feature#version GkeHubFeature#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

