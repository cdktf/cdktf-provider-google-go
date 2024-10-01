// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containernodepool


type ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfig struct {
	// Amount of 1G hugepages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/container_node_pool#hugepage_size_1g ContainerNodePool#hugepage_size_1g}
	HugepageSize1G *float64 `field:"optional" json:"hugepageSize1G" yaml:"hugepageSize1G"`
	// Amount of 2M hugepages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/container_node_pool#hugepage_size_2m ContainerNodePool#hugepage_size_2m}
	HugepageSize2M *float64 `field:"optional" json:"hugepageSize2M" yaml:"hugepageSize2M"`
}

