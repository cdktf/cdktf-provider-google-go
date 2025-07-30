// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computerouterroutepolicy


type ComputeRouterRoutePolicyTermsActions struct {
	// Textual representation of an expression in Common Expression Language syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_router_route_policy#expression ComputeRouterRoutePolicy#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Description of the expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_router_route_policy#description ComputeRouterRoutePolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_router_route_policy#location ComputeRouterRoutePolicy#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Title for the expression, i.e. a short string describing its purpose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_router_route_policy#title ComputeRouterRoutePolicy#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

