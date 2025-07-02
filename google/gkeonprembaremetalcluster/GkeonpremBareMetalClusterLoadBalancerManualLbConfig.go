// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterLoadBalancerManualLbConfig struct {
	// Whether manual load balancing is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/gkeonprem_bare_metal_cluster#enabled GkeonpremBareMetalCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

