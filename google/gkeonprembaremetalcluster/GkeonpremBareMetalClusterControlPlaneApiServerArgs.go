// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterControlPlaneApiServerArgs struct {
	// The argument name as it appears on the API Server command line please make sure to remove the leading dashes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/gkeonprem_bare_metal_cluster#argument GkeonpremBareMetalCluster#argument}
	Argument *string `field:"required" json:"argument" yaml:"argument"`
	// The value of the arg as it will be passed to the API Server command line.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/gkeonprem_bare_metal_cluster#value GkeonpremBareMetalCluster#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

