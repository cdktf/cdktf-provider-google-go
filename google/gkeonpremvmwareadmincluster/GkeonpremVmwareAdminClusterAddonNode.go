// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwareadmincluster


type GkeonpremVmwareAdminClusterAddonNode struct {
	// auto_resize_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/gkeonprem_vmware_admin_cluster#auto_resize_config GkeonpremVmwareAdminCluster#auto_resize_config}
	AutoResizeConfig *GkeonpremVmwareAdminClusterAddonNodeAutoResizeConfig `field:"optional" json:"autoResizeConfig" yaml:"autoResizeConfig"`
}

