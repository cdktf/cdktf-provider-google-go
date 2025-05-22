// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeaturemembership


type GkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainers struct {
	// The name of the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/gke_hub_feature_membership#container_name GkeHubFeatureMembership#container_name}
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The CPU limit of the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/gke_hub_feature_membership#cpu_limit GkeHubFeatureMembership#cpu_limit}
	CpuLimit *string `field:"optional" json:"cpuLimit" yaml:"cpuLimit"`
	// The CPU request of the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/gke_hub_feature_membership#cpu_request GkeHubFeatureMembership#cpu_request}
	CpuRequest *string `field:"optional" json:"cpuRequest" yaml:"cpuRequest"`
	// The memory limit of the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/gke_hub_feature_membership#memory_limit GkeHubFeatureMembership#memory_limit}
	MemoryLimit *string `field:"optional" json:"memoryLimit" yaml:"memoryLimit"`
	// The memory request of the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/gke_hub_feature_membership#memory_request GkeHubFeatureMembership#memory_request}
	MemoryRequest *string `field:"optional" json:"memoryRequest" yaml:"memoryRequest"`
}

