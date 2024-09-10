// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterLoadBalancer struct {
	// port_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/gkeonprem_bare_metal_cluster#port_config GkeonpremBareMetalCluster#port_config}
	PortConfig *GkeonpremBareMetalClusterLoadBalancerPortConfig `field:"required" json:"portConfig" yaml:"portConfig"`
	// vip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/gkeonprem_bare_metal_cluster#vip_config GkeonpremBareMetalCluster#vip_config}
	VipConfig *GkeonpremBareMetalClusterLoadBalancerVipConfig `field:"required" json:"vipConfig" yaml:"vipConfig"`
	// bgp_lb_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/gkeonprem_bare_metal_cluster#bgp_lb_config GkeonpremBareMetalCluster#bgp_lb_config}
	BgpLbConfig *GkeonpremBareMetalClusterLoadBalancerBgpLbConfig `field:"optional" json:"bgpLbConfig" yaml:"bgpLbConfig"`
	// manual_lb_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/gkeonprem_bare_metal_cluster#manual_lb_config GkeonpremBareMetalCluster#manual_lb_config}
	ManualLbConfig *GkeonpremBareMetalClusterLoadBalancerManualLbConfig `field:"optional" json:"manualLbConfig" yaml:"manualLbConfig"`
	// metal_lb_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/gkeonprem_bare_metal_cluster#metal_lb_config GkeonpremBareMetalCluster#metal_lb_config}
	MetalLbConfig *GkeonpremBareMetalClusterLoadBalancerMetalLbConfig `field:"optional" json:"metalLbConfig" yaml:"metalLbConfig"`
}

