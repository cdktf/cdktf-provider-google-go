// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkserviceshttproute


type NetworkServicesHttpRouteRulesActionFaultInjectionPolicyAbort struct {
	// The HTTP status code used to abort the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/network_services_http_route#http_status NetworkServicesHttpRoute#http_status}
	HttpStatus *float64 `field:"optional" json:"httpStatus" yaml:"httpStatus"`
	// The percentage of traffic which will be aborted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/network_services_http_route#percentage NetworkServicesHttpRoute#percentage}
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
}

