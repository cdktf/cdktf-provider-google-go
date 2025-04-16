// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicesgrpcroute


type NetworkServicesGrpcRouteRulesActionDestinations struct {
	// The URL of a BackendService to route traffic to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/network_services_grpc_route#service_name NetworkServicesGrpcRoute#service_name}
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Specifies the proportion of requests forwarded to the backend referenced by the serviceName field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/network_services_grpc_route#weight NetworkServicesGrpcRoute#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

