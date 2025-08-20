// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwareadmincluster


type GkeonpremVmwareAdminClusterNetworkConfigDhcpIpConfig struct {
	// enabled is a flag to mark if DHCP IP allocation is used for VMware admin clusters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/gkeonprem_vmware_admin_cluster#enabled GkeonpremVmwareAdminCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

