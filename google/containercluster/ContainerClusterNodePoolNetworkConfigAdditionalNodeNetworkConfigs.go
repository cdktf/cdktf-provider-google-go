// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodePoolNetworkConfigAdditionalNodeNetworkConfigs struct {
	// Name of the VPC where the additional interface belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/container_cluster#network ContainerCluster#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// Name of the subnetwork where the additional interface belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/container_cluster#subnetwork ContainerCluster#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
}

