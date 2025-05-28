// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicesgrpcroute


type NetworkServicesGrpcRouteRulesActionRetryPolicy struct {
	// Specifies the allowed number of retries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/network_services_grpc_route#num_retries NetworkServicesGrpcRoute#num_retries}
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
	// Specifies one or more conditions when this retry policy applies. Possible values: ["connect-failure", "refused-stream", "cancelled", "deadline-exceeded", "resource-exhausted", "unavailable"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/network_services_grpc_route#retry_conditions NetworkServicesGrpcRoute#retry_conditions}
	RetryConditions *[]*string `field:"optional" json:"retryConditions" yaml:"retryConditions"`
}

