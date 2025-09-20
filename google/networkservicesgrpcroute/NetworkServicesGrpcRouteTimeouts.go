// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicesgrpcroute


type NetworkServicesGrpcRouteTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/network_services_grpc_route#create NetworkServicesGrpcRoute#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/network_services_grpc_route#delete NetworkServicesGrpcRoute#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/network_services_grpc_route#update NetworkServicesGrpcRoute#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

