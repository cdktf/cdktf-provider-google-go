// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkserviceshttproute


type NetworkServicesHttpRouteRulesActionFaultInjectionPolicyDelay struct {
	// Specify a fixed delay before forwarding the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/network_services_http_route#fixed_delay NetworkServicesHttpRoute#fixed_delay}
	FixedDelay *string `field:"optional" json:"fixedDelay" yaml:"fixedDelay"`
	// The percentage of traffic on which delay will be injected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/network_services_http_route#percentage NetworkServicesHttpRoute#percentage}
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
}

