// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computerouterroutepolicy


type ComputeRouterRoutePolicyTerms struct {
	// The evaluation priority for this term, which must be between 0 (inclusive) and 231 (exclusive), and unique within the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/compute_router_route_policy#priority ComputeRouterRoutePolicy#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/compute_router_route_policy#actions ComputeRouterRoutePolicy#actions}
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/compute_router_route_policy#match ComputeRouterRoutePolicy#match}
	Match *ComputeRouterRoutePolicyTermsMatch `field:"optional" json:"match" yaml:"match"`
}

