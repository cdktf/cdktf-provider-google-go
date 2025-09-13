// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicesgrpcroute


type NetworkServicesGrpcRouteRulesMatchesMethod struct {
	// Required. Name of the method to match against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/network_services_grpc_route#grpc_method NetworkServicesGrpcRoute#grpc_method}
	GrpcMethod *string `field:"required" json:"grpcMethod" yaml:"grpcMethod"`
	// Required. Name of the service to match against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/network_services_grpc_route#grpc_service NetworkServicesGrpcRoute#grpc_service}
	GrpcService *string `field:"required" json:"grpcService" yaml:"grpcService"`
	// Specifies that matches are case sensitive. The default value is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/network_services_grpc_route#case_sensitive NetworkServicesGrpcRoute#case_sensitive}
	CaseSensitive interface{} `field:"optional" json:"caseSensitive" yaml:"caseSensitive"`
}

