// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwareadmincluster


type GkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfig struct {
	// control_plane_ip_block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/gkeonprem_vmware_admin_cluster#control_plane_ip_block GkeonpremVmwareAdminCluster#control_plane_ip_block}
	ControlPlaneIpBlock *GkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlock `field:"optional" json:"controlPlaneIpBlock" yaml:"controlPlaneIpBlock"`
}

