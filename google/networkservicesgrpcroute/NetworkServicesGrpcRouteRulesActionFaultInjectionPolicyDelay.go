// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicesgrpcroute


type NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelay struct {
	// Specify a fixed delay before forwarding the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/network_services_grpc_route#fixed_delay NetworkServicesGrpcRoute#fixed_delay}
	FixedDelay *string `field:"optional" json:"fixedDelay" yaml:"fixedDelay"`
	// The percentage of traffic on which delay will be injected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/network_services_grpc_route#percentage NetworkServicesGrpcRoute#percentage}
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
}

