// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterNetworkConfigMultipleNetworkInterfacesConfig struct {
	// Whether to enable multiple network interfaces for your pods. When set network_config.advanced_networking is automatically set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/gkeonprem_bare_metal_cluster#enabled GkeonpremBareMetalCluster#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

