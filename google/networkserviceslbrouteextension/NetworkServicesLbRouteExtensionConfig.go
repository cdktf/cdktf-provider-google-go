// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkserviceslbrouteextension

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkServicesLbRouteExtensionConfig struct {
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
	// extension_chains block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/network_services_lb_route_extension#extension_chains NetworkServicesLbRouteExtension#extension_chains}
	ExtensionChains interface{} `field:"required" json:"extensionChains" yaml:"extensionChains"`
	// A list of references to the forwarding rules to which this service extension is attached to.
	//
	// At least one forwarding rule is required. There can be only one LbRouteExtension resource per forwarding rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/network_services_lb_route_extension#forwarding_rules NetworkServicesLbRouteExtension#forwarding_rules}
	ForwardingRules *[]*string `field:"required" json:"forwardingRules" yaml:"forwardingRules"`
	// All backend services and forwarding rules referenced by this extension must share the same load balancing scheme.
	//
	// For more information, refer to [Choosing a load balancer](https://cloud.google.com/load-balancing/docs/backend-service) and
	// [Supported application load balancers](https://cloud.google.com/service-extensions/docs/callouts-overview#supported-lbs). Possible values: ["INTERNAL_MANAGED", "EXTERNAL_MANAGED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/network_services_lb_route_extension#load_balancing_scheme NetworkServicesLbRouteExtension#load_balancing_scheme}
	LoadBalancingScheme *string `field:"required" json:"loadBalancingScheme" yaml:"loadBalancingScheme"`
	// The location of the route extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/network_services_lb_route_extension#location NetworkServicesLbRouteExtension#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Name of the LbRouteExtension resource in the following format: projects/{project}/locations/{location}/lbRouteExtensions/{lbRouteExtension}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/network_services_lb_route_extension#name NetworkServicesLbRouteExtension#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A human-readable description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/network_services_lb_route_extension#description NetworkServicesLbRouteExtension#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/network_services_lb_route_extension#id NetworkServicesLbRouteExtension#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of labels associated with the LbRouteExtension resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/network_services_lb_route_extension#labels NetworkServicesLbRouteExtension#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/network_services_lb_route_extension#project NetworkServicesLbRouteExtension#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/network_services_lb_route_extension#timeouts NetworkServicesLbRouteExtension#timeouts}
	Timeouts *NetworkServicesLbRouteExtensionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

