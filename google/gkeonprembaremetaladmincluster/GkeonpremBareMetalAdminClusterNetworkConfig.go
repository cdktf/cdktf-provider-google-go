// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetaladmincluster


type GkeonpremBareMetalAdminClusterNetworkConfig struct {
	// island_mode_cidr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/gkeonprem_bare_metal_admin_cluster#island_mode_cidr GkeonpremBareMetalAdminCluster#island_mode_cidr}
	IslandModeCidr *GkeonpremBareMetalAdminClusterNetworkConfigIslandModeCidr `field:"optional" json:"islandModeCidr" yaml:"islandModeCidr"`
}

