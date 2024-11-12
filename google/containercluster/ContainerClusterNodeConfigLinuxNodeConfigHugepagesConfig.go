// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodeConfigLinuxNodeConfigHugepagesConfig struct {
	// Amount of 1G hugepages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/container_cluster#hugepage_size_1g ContainerCluster#hugepage_size_1g}
	HugepageSize1G *float64 `field:"optional" json:"hugepageSize1G" yaml:"hugepageSize1G"`
	// Amount of 2M hugepages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/container_cluster#hugepage_size_2m ContainerCluster#hugepage_size_2m}
	HugepageSize2M *float64 `field:"optional" json:"hugepageSize2M" yaml:"hugepageSize2M"`
}

