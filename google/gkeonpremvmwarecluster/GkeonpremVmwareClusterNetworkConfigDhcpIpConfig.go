// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwarecluster


type GkeonpremVmwareClusterNetworkConfigDhcpIpConfig struct {
	// enabled is a flag to mark if DHCP IP allocation is used for VMware user clusters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/gkeonprem_vmware_cluster#enabled GkeonpremVmwareCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

