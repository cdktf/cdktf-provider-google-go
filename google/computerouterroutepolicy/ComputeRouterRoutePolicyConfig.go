// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computerouterroutepolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRouterRoutePolicyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Name of the route policy.
	//
	// This policy's name, which must be a resource ID segment and unique within all policies owned by the Router
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/compute_router_route_policy#name ComputeRouterRoutePolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the Cloud Router in which this route policy will be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/compute_router_route_policy#router ComputeRouterRoutePolicy#router}
	Router *string `field:"required" json:"router" yaml:"router"`
	// terms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/compute_router_route_policy#terms ComputeRouterRoutePolicy#terms}
	Terms interface{} `field:"required" json:"terms" yaml:"terms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/compute_router_route_policy#id ComputeRouterRoutePolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/compute_router_route_policy#project ComputeRouterRoutePolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Region where the router and NAT reside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/compute_router_route_policy#region ComputeRouterRoutePolicy#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/compute_router_route_policy#timeouts ComputeRouterRoutePolicy#timeouts}
	Timeouts *ComputeRouterRoutePolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// This is policy's type, which is one of IMPORT or EXPORT Possible values: ["ROUTE_POLICY_TYPE_IMPORT", "ROUTE_POLICY_TYPE_EXPORT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/compute_router_route_policy#type ComputeRouterRoutePolicy#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

