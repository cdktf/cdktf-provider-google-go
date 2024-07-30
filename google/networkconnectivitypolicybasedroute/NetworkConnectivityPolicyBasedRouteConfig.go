// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkconnectivitypolicybasedroute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkConnectivityPolicyBasedRouteConfig struct {
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
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/network_connectivity_policy_based_route#filter NetworkConnectivityPolicyBasedRoute#filter}
	Filter *NetworkConnectivityPolicyBasedRouteFilter `field:"required" json:"filter" yaml:"filter"`
	// The name of the policy based route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/network_connectivity_policy_based_route#name NetworkConnectivityPolicyBasedRoute#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Fully-qualified URL of the network that this route applies to, for example: projects/my-project/global/networks/my-network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/network_connectivity_policy_based_route#network NetworkConnectivityPolicyBasedRoute#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/network_connectivity_policy_based_route#description NetworkConnectivityPolicyBasedRoute#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/network_connectivity_policy_based_route#id NetworkConnectivityPolicyBasedRoute#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// interconnect_attachment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/network_connectivity_policy_based_route#interconnect_attachment NetworkConnectivityPolicyBasedRoute#interconnect_attachment}
	InterconnectAttachment *NetworkConnectivityPolicyBasedRouteInterconnectAttachment `field:"optional" json:"interconnectAttachment" yaml:"interconnectAttachment"`
	// User-defined labels.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/network_connectivity_policy_based_route#labels NetworkConnectivityPolicyBasedRoute#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The IP address of a global-access-enabled L4 ILB that is the next hop for matching packets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/network_connectivity_policy_based_route#next_hop_ilb_ip NetworkConnectivityPolicyBasedRoute#next_hop_ilb_ip}
	NextHopIlbIp *string `field:"optional" json:"nextHopIlbIp" yaml:"nextHopIlbIp"`
	// Other routes that will be referenced to determine the next hop of the packet. Possible values: ["DEFAULT_ROUTING"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/network_connectivity_policy_based_route#next_hop_other_routes NetworkConnectivityPolicyBasedRoute#next_hop_other_routes}
	NextHopOtherRoutes *string `field:"optional" json:"nextHopOtherRoutes" yaml:"nextHopOtherRoutes"`
	// The priority of this policy-based route.
	//
	// Priority is used to break ties in cases where there are more than one matching policy-based routes found. In cases where multiple policy-based routes are matched, the one with the lowest-numbered priority value wins. The default value is 1000. The priority value must be from 1 to 65535, inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/network_connectivity_policy_based_route#priority NetworkConnectivityPolicyBasedRoute#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/network_connectivity_policy_based_route#project NetworkConnectivityPolicyBasedRoute#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/network_connectivity_policy_based_route#timeouts NetworkConnectivityPolicyBasedRoute#timeouts}
	Timeouts *NetworkConnectivityPolicyBasedRouteTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// virtual_machine block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/network_connectivity_policy_based_route#virtual_machine NetworkConnectivityPolicyBasedRoute#virtual_machine}
	VirtualMachine *NetworkConnectivityPolicyBasedRouteVirtualMachine `field:"optional" json:"virtualMachine" yaml:"virtualMachine"`
}

