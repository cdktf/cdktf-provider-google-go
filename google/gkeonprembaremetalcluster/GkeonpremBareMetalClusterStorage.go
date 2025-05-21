// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterStorage struct {
	// lvp_node_mounts_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/gkeonprem_bare_metal_cluster#lvp_node_mounts_config GkeonpremBareMetalCluster#lvp_node_mounts_config}
	LvpNodeMountsConfig *GkeonpremBareMetalClusterStorageLvpNodeMountsConfig `field:"required" json:"lvpNodeMountsConfig" yaml:"lvpNodeMountsConfig"`
	// lvp_share_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/gkeonprem_bare_metal_cluster#lvp_share_config GkeonpremBareMetalCluster#lvp_share_config}
	LvpShareConfig *GkeonpremBareMetalClusterStorageLvpShareConfig `field:"required" json:"lvpShareConfig" yaml:"lvpShareConfig"`
}

