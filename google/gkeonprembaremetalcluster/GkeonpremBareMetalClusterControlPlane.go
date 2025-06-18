// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterControlPlane struct {
	// control_plane_node_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/gkeonprem_bare_metal_cluster#control_plane_node_pool_config GkeonpremBareMetalCluster#control_plane_node_pool_config}
	ControlPlaneNodePoolConfig *GkeonpremBareMetalClusterControlPlaneControlPlaneNodePoolConfig `field:"required" json:"controlPlaneNodePoolConfig" yaml:"controlPlaneNodePoolConfig"`
	// api_server_args block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/gkeonprem_bare_metal_cluster#api_server_args GkeonpremBareMetalCluster#api_server_args}
	ApiServerArgs interface{} `field:"optional" json:"apiServerArgs" yaml:"apiServerArgs"`
}

