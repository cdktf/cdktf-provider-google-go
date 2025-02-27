// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodePoolAutoConfig struct {
	// linux_node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/container_cluster#linux_node_config ContainerCluster#linux_node_config}
	LinuxNodeConfig *ContainerClusterNodePoolAutoConfigLinuxNodeConfig `field:"optional" json:"linuxNodeConfig" yaml:"linuxNodeConfig"`
	// network_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/container_cluster#network_tags ContainerCluster#network_tags}
	NetworkTags *ContainerClusterNodePoolAutoConfigNetworkTags `field:"optional" json:"networkTags" yaml:"networkTags"`
	// node_kubelet_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/container_cluster#node_kubelet_config ContainerCluster#node_kubelet_config}
	NodeKubeletConfig *ContainerClusterNodePoolAutoConfigNodeKubeletConfig `field:"optional" json:"nodeKubeletConfig" yaml:"nodeKubeletConfig"`
	// A map of resource manager tags.
	//
	// Resource manager tag keys and values have the same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored (both PUT & PATCH) when empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/container_cluster#resource_manager_tags ContainerCluster#resource_manager_tags}
	ResourceManagerTags *map[string]*string `field:"optional" json:"resourceManagerTags" yaml:"resourceManagerTags"`
}

