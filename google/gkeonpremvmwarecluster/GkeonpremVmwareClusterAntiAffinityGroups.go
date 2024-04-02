// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwarecluster


type GkeonpremVmwareClusterAntiAffinityGroups struct {
	// Spread nodes across at least three physical hosts (requires at least three hosts). Enabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.23.0/docs/resources/gkeonprem_vmware_cluster#aag_config_disabled GkeonpremVmwareCluster#aag_config_disabled}
	AagConfigDisabled interface{} `field:"required" json:"aagConfigDisabled" yaml:"aagConfigDisabled"`
}

