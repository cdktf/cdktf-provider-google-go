// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkserviceshttproute


type NetworkServicesHttpRouteRulesActionRequestMirrorPolicy struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/network_services_http_route#destination NetworkServicesHttpRoute#destination}
	Destination *NetworkServicesHttpRouteRulesActionRequestMirrorPolicyDestination `field:"optional" json:"destination" yaml:"destination"`
}

