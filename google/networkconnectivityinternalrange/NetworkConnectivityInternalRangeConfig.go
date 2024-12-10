// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkconnectivityinternalrange

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkConnectivityInternalRangeConfig struct {
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
	// The name of the policy based route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/network_connectivity_internal_range#name NetworkConnectivityInternalRange#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Fully-qualified URL of the network that this route applies to, for example: projects/my-project/global/networks/my-network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/network_connectivity_internal_range#network NetworkConnectivityInternalRange#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// The type of peering set for this internal range. Possible values: ["FOR_SELF", "FOR_PEER", "NOT_SHARED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/network_connectivity_internal_range#peering NetworkConnectivityInternalRange#peering}
	Peering *string `field:"required" json:"peering" yaml:"peering"`
	// The type of usage set for this InternalRange. Possible values: ["FOR_VPC", "EXTERNAL_TO_VPC", "FOR_MIGRATION"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/network_connectivity_internal_range#usage NetworkConnectivityInternalRange#usage}
	Usage *string `field:"required" json:"usage" yaml:"usage"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/network_connectivity_internal_range#description NetworkConnectivityInternalRange#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/network_connectivity_internal_range#id NetworkConnectivityInternalRange#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The IP range that this internal range defines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/network_connectivity_internal_range#ip_cidr_range NetworkConnectivityInternalRange#ip_cidr_range}
	IpCidrRange *string `field:"optional" json:"ipCidrRange" yaml:"ipCidrRange"`
	// User-defined labels.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/network_connectivity_internal_range#labels NetworkConnectivityInternalRange#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// migration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/network_connectivity_internal_range#migration NetworkConnectivityInternalRange#migration}
	Migration *NetworkConnectivityInternalRangeMigration `field:"optional" json:"migration" yaml:"migration"`
	// Optional. Types of resources that are allowed to overlap with the current internal range. Possible values: ["OVERLAP_ROUTE_RANGE", "OVERLAP_EXISTING_SUBNET_RANGE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/network_connectivity_internal_range#overlaps NetworkConnectivityInternalRange#overlaps}
	Overlaps *[]*string `field:"optional" json:"overlaps" yaml:"overlaps"`
	// An alternate to ipCidrRange.
	//
	// Can be set when trying to create a reservation that automatically finds a free range of the given size.
	// If both ipCidrRange and prefixLength are set, there is an error if the range sizes do not match. Can also be used during updates to change the range size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/network_connectivity_internal_range#prefix_length NetworkConnectivityInternalRange#prefix_length}
	PrefixLength *float64 `field:"optional" json:"prefixLength" yaml:"prefixLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/network_connectivity_internal_range#project NetworkConnectivityInternalRange#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Optional.
	//
	// Can be set to narrow down or pick a different address space while searching for a free range.
	// If not set, defaults to the "10.0.0.0/8" address space. This can be used to search in other rfc-1918 address spaces like "172.16.0.0/12" and "192.168.0.0/16" or non-rfc-1918 address spaces used in the VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/network_connectivity_internal_range#target_cidr_range NetworkConnectivityInternalRange#target_cidr_range}
	TargetCidrRange *[]*string `field:"optional" json:"targetCidrRange" yaml:"targetCidrRange"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/network_connectivity_internal_range#timeouts NetworkConnectivityInternalRange#timeouts}
	Timeouts *NetworkConnectivityInternalRangeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

