// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkserviceslbrouteextension


type NetworkServicesLbRouteExtensionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/network_services_lb_route_extension#create NetworkServicesLbRouteExtension#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/network_services_lb_route_extension#delete NetworkServicesLbRouteExtension#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/network_services_lb_route_extension#update NetworkServicesLbRouteExtension#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

