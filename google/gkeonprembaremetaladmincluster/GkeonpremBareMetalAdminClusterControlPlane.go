// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetaladmincluster


type GkeonpremBareMetalAdminClusterControlPlane struct {
	// control_plane_node_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/gkeonprem_bare_metal_admin_cluster#control_plane_node_pool_config GkeonpremBareMetalAdminCluster#control_plane_node_pool_config}
	ControlPlaneNodePoolConfig *GkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfig `field:"required" json:"controlPlaneNodePoolConfig" yaml:"controlPlaneNodePoolConfig"`
	// api_server_args block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/gkeonprem_bare_metal_admin_cluster#api_server_args GkeonpremBareMetalAdminCluster#api_server_args}
	ApiServerArgs interface{} `field:"optional" json:"apiServerArgs" yaml:"apiServerArgs"`
}

