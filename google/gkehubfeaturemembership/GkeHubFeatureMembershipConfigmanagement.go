// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeaturemembership


type GkeHubFeatureMembershipConfigmanagement struct {
	// binauthz block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/gke_hub_feature_membership#binauthz GkeHubFeatureMembership#binauthz}
	Binauthz *GkeHubFeatureMembershipConfigmanagementBinauthz `field:"optional" json:"binauthz" yaml:"binauthz"`
	// config_sync block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/gke_hub_feature_membership#config_sync GkeHubFeatureMembership#config_sync}
	ConfigSync *GkeHubFeatureMembershipConfigmanagementConfigSync `field:"optional" json:"configSync" yaml:"configSync"`
	// hierarchy_controller block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/gke_hub_feature_membership#hierarchy_controller GkeHubFeatureMembership#hierarchy_controller}
	HierarchyController *GkeHubFeatureMembershipConfigmanagementHierarchyController `field:"optional" json:"hierarchyController" yaml:"hierarchyController"`
	// Set this field to MANAGEMENT_AUTOMATIC to enable Config Sync auto-upgrades, and set this field to MANAGEMENT_MANUAL or MANAGEMENT_UNSPECIFIED to disable Config Sync auto-upgrades.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/gke_hub_feature_membership#management GkeHubFeatureMembership#management}
	Management *string `field:"optional" json:"management" yaml:"management"`
	// policy_controller block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/gke_hub_feature_membership#policy_controller GkeHubFeatureMembership#policy_controller}
	PolicyController *GkeHubFeatureMembershipConfigmanagementPolicyController `field:"optional" json:"policyController" yaml:"policyController"`
	// Optional. Version of ACM to install. Defaults to the latest version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/gke_hub_feature_membership#version GkeHubFeatureMembership#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

