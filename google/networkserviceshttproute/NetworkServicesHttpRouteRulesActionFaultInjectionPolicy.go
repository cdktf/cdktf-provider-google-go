// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkserviceshttproute


type NetworkServicesHttpRouteRulesActionFaultInjectionPolicy struct {
	// abort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/network_services_http_route#abort NetworkServicesHttpRoute#abort}
	Abort *NetworkServicesHttpRouteRulesActionFaultInjectionPolicyAbort `field:"optional" json:"abort" yaml:"abort"`
	// delay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/network_services_http_route#delay NetworkServicesHttpRoute#delay}
	Delay *NetworkServicesHttpRouteRulesActionFaultInjectionPolicyDelay `field:"optional" json:"delay" yaml:"delay"`
}

