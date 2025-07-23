// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containernodepool


type ContainerNodePoolNetworkConfigAdditionalNodeNetworkConfigs struct {
	// Name of the VPC where the additional interface belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_node_pool#network ContainerNodePool#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// Name of the subnetwork where the additional interface belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_node_pool#subnetwork ContainerNodePool#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
}

