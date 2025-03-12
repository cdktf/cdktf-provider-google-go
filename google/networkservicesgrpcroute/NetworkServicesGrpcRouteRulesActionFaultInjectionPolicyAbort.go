// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicesgrpcroute


type NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbort struct {
	// The HTTP status code used to abort the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/network_services_grpc_route#http_status NetworkServicesGrpcRoute#http_status}
	HttpStatus *float64 `field:"optional" json:"httpStatus" yaml:"httpStatus"`
	// The percentage of traffic which will be aborted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/network_services_grpc_route#percentage NetworkServicesGrpcRoute#percentage}
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
}

