// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetaladmincluster


type GkeonpremBareMetalAdminClusterLoadBalancerManualLbConfig struct {
	// Whether manual load balancing is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/gkeonprem_bare_metal_admin_cluster#enabled GkeonpremBareMetalAdminCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

