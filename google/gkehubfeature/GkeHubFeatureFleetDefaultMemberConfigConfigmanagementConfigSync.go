// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeature


type GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSync struct {
	// git block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.16.0/docs/resources/gke_hub_feature#git GkeHubFeature#git}
	Git *GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGit `field:"optional" json:"git" yaml:"git"`
	// oci block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.16.0/docs/resources/gke_hub_feature#oci GkeHubFeature#oci}
	Oci *GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOci `field:"optional" json:"oci" yaml:"oci"`
	// Specifies whether the Config Sync Repo is in hierarchical or unstructured mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.16.0/docs/resources/gke_hub_feature#source_format GkeHubFeature#source_format}
	SourceFormat *string `field:"optional" json:"sourceFormat" yaml:"sourceFormat"`
}

