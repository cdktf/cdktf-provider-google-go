// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwarecluster


type GkeonpremVmwareClusterLoadBalancerMetalLbConfig struct {
	// address_pools block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/gkeonprem_vmware_cluster#address_pools GkeonpremVmwareCluster#address_pools}
	AddressPools interface{} `field:"required" json:"addressPools" yaml:"addressPools"`
}

