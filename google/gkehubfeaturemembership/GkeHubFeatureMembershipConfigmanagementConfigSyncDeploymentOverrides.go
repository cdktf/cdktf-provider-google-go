// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeaturemembership


type GkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverrides struct {
	// containers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/gke_hub_feature_membership#containers GkeHubFeatureMembership#containers}
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// The name of the Deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/gke_hub_feature_membership#deployment_name GkeHubFeatureMembership#deployment_name}
	DeploymentName *string `field:"optional" json:"deploymentName" yaml:"deploymentName"`
	// The namespace of the Deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/gke_hub_feature_membership#deployment_namespace GkeHubFeatureMembership#deployment_namespace}
	DeploymentNamespace *string `field:"optional" json:"deploymentNamespace" yaml:"deploymentNamespace"`
}

