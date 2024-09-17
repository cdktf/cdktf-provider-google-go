// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterLoadBalancerBgpLbConfig struct {
	// address_pools block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.3.0/docs/resources/gkeonprem_bare_metal_cluster#address_pools GkeonpremBareMetalCluster#address_pools}
	AddressPools interface{} `field:"required" json:"addressPools" yaml:"addressPools"`
	// BGP autonomous system number (ASN) of the cluster. This field can be updated after cluster creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.3.0/docs/resources/gkeonprem_bare_metal_cluster#asn GkeonpremBareMetalCluster#asn}
	Asn *float64 `field:"required" json:"asn" yaml:"asn"`
	// bgp_peer_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.3.0/docs/resources/gkeonprem_bare_metal_cluster#bgp_peer_configs GkeonpremBareMetalCluster#bgp_peer_configs}
	BgpPeerConfigs interface{} `field:"required" json:"bgpPeerConfigs" yaml:"bgpPeerConfigs"`
	// load_balancer_node_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.3.0/docs/resources/gkeonprem_bare_metal_cluster#load_balancer_node_pool_config GkeonpremBareMetalCluster#load_balancer_node_pool_config}
	LoadBalancerNodePoolConfig *GkeonpremBareMetalClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfig `field:"optional" json:"loadBalancerNodePoolConfig" yaml:"loadBalancerNodePoolConfig"`
}

