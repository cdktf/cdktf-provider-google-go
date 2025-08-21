// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicestcproute


type NetworkServicesTcpRouteRulesAction struct {
	// destinations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/network_services_tcp_route#destinations NetworkServicesTcpRoute#destinations}
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// Specifies the idle timeout for the selected route.
	//
	// The idle timeout is defined as the period in which there are no bytes sent or received on either the upstream or downstream connection. If not set, the default idle timeout is 30 seconds. If set to 0s, the timeout will be disabled.
	//
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/network_services_tcp_route#idle_timeout NetworkServicesTcpRoute#idle_timeout}
	IdleTimeout *string `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// If true, Router will use the destination IP and port of the original connection as the destination of the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/network_services_tcp_route#original_destination NetworkServicesTcpRoute#original_destination}
	OriginalDestination interface{} `field:"optional" json:"originalDestination" yaml:"originalDestination"`
}

