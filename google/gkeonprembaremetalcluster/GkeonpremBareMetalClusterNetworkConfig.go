// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterNetworkConfig struct {
	// Enables the use of advanced Anthos networking features, such as Bundled Load Balancing with BGP or the egress NAT gateway.
	//
	// Setting configuration for advanced networking features will automatically
	// set this flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/gkeonprem_bare_metal_cluster#advanced_networking GkeonpremBareMetalCluster#advanced_networking}
	AdvancedNetworking interface{} `field:"optional" json:"advancedNetworking" yaml:"advancedNetworking"`
	// island_mode_cidr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/gkeonprem_bare_metal_cluster#island_mode_cidr GkeonpremBareMetalCluster#island_mode_cidr}
	IslandModeCidr *GkeonpremBareMetalClusterNetworkConfigIslandModeCidr `field:"optional" json:"islandModeCidr" yaml:"islandModeCidr"`
	// multiple_network_interfaces_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/gkeonprem_bare_metal_cluster#multiple_network_interfaces_config GkeonpremBareMetalCluster#multiple_network_interfaces_config}
	MultipleNetworkInterfacesConfig *GkeonpremBareMetalClusterNetworkConfigMultipleNetworkInterfacesConfig `field:"optional" json:"multipleNetworkInterfacesConfig" yaml:"multipleNetworkInterfacesConfig"`
	// sr_iov_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/gkeonprem_bare_metal_cluster#sr_iov_config GkeonpremBareMetalCluster#sr_iov_config}
	SrIovConfig *GkeonpremBareMetalClusterNetworkConfigSrIovConfig `field:"optional" json:"srIovConfig" yaml:"srIovConfig"`
}

