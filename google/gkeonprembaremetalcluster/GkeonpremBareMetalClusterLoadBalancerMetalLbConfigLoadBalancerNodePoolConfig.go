// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfig struct {
	// node_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/gkeonprem_bare_metal_cluster#node_pool_config GkeonpremBareMetalCluster#node_pool_config}
	NodePoolConfig *GkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfig `field:"optional" json:"nodePoolConfig" yaml:"nodePoolConfig"`
}

