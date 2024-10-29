// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicesedgecacheservice


type NetworkServicesEdgeCacheServiceRouting struct {
	// host_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/network_services_edge_cache_service#host_rule NetworkServicesEdgeCacheService#host_rule}
	HostRule interface{} `field:"required" json:"hostRule" yaml:"hostRule"`
	// path_matcher block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/network_services_edge_cache_service#path_matcher NetworkServicesEdgeCacheService#path_matcher}
	PathMatcher interface{} `field:"required" json:"pathMatcher" yaml:"pathMatcher"`
}

