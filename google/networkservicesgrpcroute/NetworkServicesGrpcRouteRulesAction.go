// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicesgrpcroute


type NetworkServicesGrpcRouteRulesAction struct {
	// destinations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/network_services_grpc_route#destinations NetworkServicesGrpcRoute#destinations}
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// fault_injection_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/network_services_grpc_route#fault_injection_policy NetworkServicesGrpcRoute#fault_injection_policy}
	FaultInjectionPolicy *NetworkServicesGrpcRouteRulesActionFaultInjectionPolicy `field:"optional" json:"faultInjectionPolicy" yaml:"faultInjectionPolicy"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/network_services_grpc_route#retry_policy NetworkServicesGrpcRoute#retry_policy}
	RetryPolicy *NetworkServicesGrpcRouteRulesActionRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// Specifies the timeout for selected route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/network_services_grpc_route#timeout NetworkServicesGrpcRoute#timeout}
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

