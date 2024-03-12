// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeaturemembership


type GkeHubFeatureMembershipConfigmanagementHierarchyController struct {
	// Whether Hierarchy Controller is enabled in this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.20.0/docs/resources/gke_hub_feature_membership#enabled GkeHubFeatureMembership#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Whether hierarchical resource quota is enabled in this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.20.0/docs/resources/gke_hub_feature_membership#enable_hierarchical_resource_quota GkeHubFeatureMembership#enable_hierarchical_resource_quota}
	EnableHierarchicalResourceQuota interface{} `field:"optional" json:"enableHierarchicalResourceQuota" yaml:"enableHierarchicalResourceQuota"`
	// Whether pod tree labels are enabled in this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.20.0/docs/resources/gke_hub_feature_membership#enable_pod_tree_labels GkeHubFeatureMembership#enable_pod_tree_labels}
	EnablePodTreeLabels interface{} `field:"optional" json:"enablePodTreeLabels" yaml:"enablePodTreeLabels"`
}

