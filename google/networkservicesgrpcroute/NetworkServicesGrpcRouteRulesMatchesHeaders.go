// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicesgrpcroute


type NetworkServicesGrpcRouteRulesMatchesHeaders struct {
	// Required. The key of the header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/network_services_grpc_route#key NetworkServicesGrpcRoute#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Required. The value of the header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/network_services_grpc_route#value NetworkServicesGrpcRoute#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// The type of match. Default value: "EXACT" Possible values: ["TYPE_UNSPECIFIED", "EXACT", "REGULAR_EXPRESSION"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/network_services_grpc_route#type NetworkServicesGrpcRoute#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

