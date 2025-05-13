// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwarecluster


type GkeonpremVmwareClusterDataplaneV2 struct {
	// Enable advanced networking which requires dataplane_v2_enabled to be set true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/gkeonprem_vmware_cluster#advanced_networking GkeonpremVmwareCluster#advanced_networking}
	AdvancedNetworking interface{} `field:"optional" json:"advancedNetworking" yaml:"advancedNetworking"`
	// Enables Dataplane V2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/gkeonprem_vmware_cluster#dataplane_v2_enabled GkeonpremVmwareCluster#dataplane_v2_enabled}
	DataplaneV2Enabled interface{} `field:"optional" json:"dataplaneV2Enabled" yaml:"dataplaneV2Enabled"`
	// Enable Dataplane V2 for clusters with Windows nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/gkeonprem_vmware_cluster#windows_dataplane_v2_enabled GkeonpremVmwareCluster#windows_dataplane_v2_enabled}
	WindowsDataplaneV2Enabled interface{} `field:"optional" json:"windowsDataplaneV2Enabled" yaml:"windowsDataplaneV2Enabled"`
}

