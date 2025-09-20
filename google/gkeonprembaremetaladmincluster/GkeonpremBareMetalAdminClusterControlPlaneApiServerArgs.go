// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetaladmincluster


type GkeonpremBareMetalAdminClusterControlPlaneApiServerArgs struct {
	// The argument name as it appears on the API Server command line please make sure to remove the leading dashes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/gkeonprem_bare_metal_admin_cluster#argument GkeonpremBareMetalAdminCluster#argument}
	Argument *string `field:"required" json:"argument" yaml:"argument"`
	// The value of the arg as it will be passed to the API Server command line.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/gkeonprem_bare_metal_admin_cluster#value GkeonpremBareMetalAdminCluster#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

