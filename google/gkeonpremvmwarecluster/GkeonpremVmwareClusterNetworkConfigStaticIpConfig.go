// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwarecluster


type GkeonpremVmwareClusterNetworkConfigStaticIpConfig struct {
	// ip_blocks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/gkeonprem_vmware_cluster#ip_blocks GkeonpremVmwareCluster#ip_blocks}
	IpBlocks interface{} `field:"required" json:"ipBlocks" yaml:"ipBlocks"`
}

