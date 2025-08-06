// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwareadmincluster


type GkeonpremVmwareAdminClusterNetworkConfigStaticIpConfigIpBlocksIps struct {
	// IP could be an IP address (like 1.2.3.4) or a CIDR (like 1.2.3.0/24).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/gkeonprem_vmware_admin_cluster#ip GkeonpremVmwareAdminCluster#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// Hostname of the machine. VM's name will be used if this field is empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/gkeonprem_vmware_admin_cluster#hostname GkeonpremVmwareAdminCluster#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
}

