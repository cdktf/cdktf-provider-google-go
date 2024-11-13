// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterLoadBalancerMetalLbConfig struct {
	// address_pools block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/gkeonprem_bare_metal_cluster#address_pools GkeonpremBareMetalCluster#address_pools}
	AddressPools interface{} `field:"required" json:"addressPools" yaml:"addressPools"`
	// load_balancer_node_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/gkeonprem_bare_metal_cluster#load_balancer_node_pool_config GkeonpremBareMetalCluster#load_balancer_node_pool_config}
	LoadBalancerNodePoolConfig *GkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfig `field:"optional" json:"loadBalancerNodePoolConfig" yaml:"loadBalancerNodePoolConfig"`
}

