// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodePoolAutoConfigNetworkTags struct {
	// List of network tags applied to auto-provisioned node pools.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/container_cluster#tags ContainerCluster#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

