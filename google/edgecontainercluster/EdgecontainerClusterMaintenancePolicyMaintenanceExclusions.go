// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgecontainercluster


type EdgecontainerClusterMaintenancePolicyMaintenanceExclusions struct {
	// A unique (per cluster) id for the window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/edgecontainer_cluster#id EdgecontainerCluster#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/edgecontainer_cluster#window EdgecontainerCluster#window}
	Window *EdgecontainerClusterMaintenancePolicyMaintenanceExclusionsWindow `field:"optional" json:"window" yaml:"window"`
}

