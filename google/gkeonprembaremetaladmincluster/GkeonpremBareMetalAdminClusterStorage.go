// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetaladmincluster


type GkeonpremBareMetalAdminClusterStorage struct {
	// lvp_node_mounts_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/gkeonprem_bare_metal_admin_cluster#lvp_node_mounts_config GkeonpremBareMetalAdminCluster#lvp_node_mounts_config}
	LvpNodeMountsConfig *GkeonpremBareMetalAdminClusterStorageLvpNodeMountsConfig `field:"required" json:"lvpNodeMountsConfig" yaml:"lvpNodeMountsConfig"`
	// lvp_share_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/gkeonprem_bare_metal_admin_cluster#lvp_share_config GkeonpremBareMetalAdminCluster#lvp_share_config}
	LvpShareConfig *GkeonpremBareMetalAdminClusterStorageLvpShareConfig `field:"required" json:"lvpShareConfig" yaml:"lvpShareConfig"`
}

