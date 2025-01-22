// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containernodepool


type ContainerNodePoolNetworkConfigAdditionalPodNetworkConfigs struct {
	// The maximum number of pods per node which use this pod network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/container_node_pool#max_pods_per_node ContainerNodePool#max_pods_per_node}
	MaxPodsPerNode *float64 `field:"optional" json:"maxPodsPerNode" yaml:"maxPodsPerNode"`
	// The name of the secondary range on the subnet which provides IP address for this pod range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/container_node_pool#secondary_pod_range ContainerNodePool#secondary_pod_range}
	SecondaryPodRange *string `field:"optional" json:"secondaryPodRange" yaml:"secondaryPodRange"`
	// Name of the subnetwork where the additional pod network belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/container_node_pool#subnetwork ContainerNodePool#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
}

