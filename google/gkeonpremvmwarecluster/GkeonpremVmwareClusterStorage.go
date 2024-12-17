// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwarecluster


type GkeonpremVmwareClusterStorage struct {
	// Whether or not to deploy vSphere CSI components in the VMware User Cluster. Enabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/gkeonprem_vmware_cluster#vsphere_csi_disabled GkeonpremVmwareCluster#vsphere_csi_disabled}
	VsphereCsiDisabled interface{} `field:"required" json:"vsphereCsiDisabled" yaml:"vsphereCsiDisabled"`
}

