// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwareadmincluster


type GkeonpremVmwareAdminClusterAddonNodeAutoResizeConfig struct {
	// Whether to enable controle plane node auto resizing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/gkeonprem_vmware_admin_cluster#enabled GkeonpremVmwareAdminCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

