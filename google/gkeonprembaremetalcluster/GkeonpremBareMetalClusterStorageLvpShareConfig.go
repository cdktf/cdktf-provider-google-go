// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterStorageLvpShareConfig struct {
	// lvp_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/gkeonprem_bare_metal_cluster#lvp_config GkeonpremBareMetalCluster#lvp_config}
	LvpConfig *GkeonpremBareMetalClusterStorageLvpShareConfigLvpConfig `field:"required" json:"lvpConfig" yaml:"lvpConfig"`
	// The number of subdirectories to create under path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/gkeonprem_bare_metal_cluster#shared_path_pv_count GkeonpremBareMetalCluster#shared_path_pv_count}
	SharedPathPvCount *float64 `field:"optional" json:"sharedPathPvCount" yaml:"sharedPathPvCount"`
}

