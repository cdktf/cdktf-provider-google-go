// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkserviceslbrouteextension


type NetworkServicesLbRouteExtensionExtensionChains struct {
	// extensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/network_services_lb_route_extension#extensions NetworkServicesLbRouteExtension#extensions}
	Extensions interface{} `field:"required" json:"extensions" yaml:"extensions"`
	// match_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/network_services_lb_route_extension#match_condition NetworkServicesLbRouteExtension#match_condition}
	MatchCondition *NetworkServicesLbRouteExtensionExtensionChainsMatchCondition `field:"required" json:"matchCondition" yaml:"matchCondition"`
	// The name for this extension chain.
	//
	// The name is logged as part of the HTTP request logs.
	// The name must conform with RFC-1034, is restricted to lower-cased letters, numbers and hyphens,
	// and can have a maximum length of 63 characters. Additionally, the first character must be a letter
	// and the last character must be a letter or a number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/network_services_lb_route_extension#name NetworkServicesLbRouteExtension#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

