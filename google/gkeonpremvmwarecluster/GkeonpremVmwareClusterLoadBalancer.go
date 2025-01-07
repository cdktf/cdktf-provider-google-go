// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwarecluster


type GkeonpremVmwareClusterLoadBalancer struct {
	// f5_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/gkeonprem_vmware_cluster#f5_config GkeonpremVmwareCluster#f5_config}
	F5Config *GkeonpremVmwareClusterLoadBalancerF5Config `field:"optional" json:"f5Config" yaml:"f5Config"`
	// manual_lb_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/gkeonprem_vmware_cluster#manual_lb_config GkeonpremVmwareCluster#manual_lb_config}
	ManualLbConfig *GkeonpremVmwareClusterLoadBalancerManualLbConfig `field:"optional" json:"manualLbConfig" yaml:"manualLbConfig"`
	// metal_lb_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/gkeonprem_vmware_cluster#metal_lb_config GkeonpremVmwareCluster#metal_lb_config}
	MetalLbConfig *GkeonpremVmwareClusterLoadBalancerMetalLbConfig `field:"optional" json:"metalLbConfig" yaml:"metalLbConfig"`
	// vip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/gkeonprem_vmware_cluster#vip_config GkeonpremVmwareCluster#vip_config}
	VipConfig *GkeonpremVmwareClusterLoadBalancerVipConfig `field:"optional" json:"vipConfig" yaml:"vipConfig"`
}

