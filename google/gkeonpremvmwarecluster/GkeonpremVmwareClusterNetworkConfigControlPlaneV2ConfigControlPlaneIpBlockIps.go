// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwarecluster


type GkeonpremVmwareClusterNetworkConfigControlPlaneV2ConfigControlPlaneIpBlockIps struct {
	// Hostname of the machine. VM's name will be used if this field is empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/gkeonprem_vmware_cluster#hostname GkeonpremVmwareCluster#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// IP could be an IP address (like 1.2.3.4) or a CIDR (like 1.2.3.0/24).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/gkeonprem_vmware_cluster#ip GkeonpremVmwareCluster#ip}
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
}

