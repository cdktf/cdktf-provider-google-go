// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodePoolDefaults struct {
	// node_config_defaults block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/container_cluster#node_config_defaults ContainerCluster#node_config_defaults}
	NodeConfigDefaults *ContainerClusterNodePoolDefaultsNodeConfigDefaults `field:"optional" json:"nodeConfigDefaults" yaml:"nodeConfigDefaults"`
}

