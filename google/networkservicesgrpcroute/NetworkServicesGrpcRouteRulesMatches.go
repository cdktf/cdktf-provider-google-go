// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicesgrpcroute


type NetworkServicesGrpcRouteRulesMatches struct {
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/network_services_grpc_route#headers NetworkServicesGrpcRoute#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/network_services_grpc_route#method NetworkServicesGrpcRoute#method}
	Method *NetworkServicesGrpcRouteRulesMatchesMethod `field:"optional" json:"method" yaml:"method"`
}

