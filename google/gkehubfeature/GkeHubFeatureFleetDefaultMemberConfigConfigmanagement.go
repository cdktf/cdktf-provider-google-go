// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeature


type GkeHubFeatureFleetDefaultMemberConfigConfigmanagement struct {
	// config_sync block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/gke_hub_feature#config_sync GkeHubFeature#config_sync}
	ConfigSync *GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSync `field:"optional" json:"configSync" yaml:"configSync"`
}

