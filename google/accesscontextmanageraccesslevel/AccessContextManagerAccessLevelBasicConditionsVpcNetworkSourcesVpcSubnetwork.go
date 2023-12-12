// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanageraccesslevel


type AccessContextManagerAccessLevelBasicConditionsVpcNetworkSourcesVpcSubnetwork struct {
	// Required.
	//
	// Network name to be allowed by this Access Level. Networks of foreign organizations requires 'compute.network.get' permission to be granted to caller.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/access_context_manager_access_level#network AccessContextManagerAccessLevel#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// CIDR block IP subnetwork specification. Must be IPv4.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/access_context_manager_access_level#vpc_ip_subnetworks AccessContextManagerAccessLevel#vpc_ip_subnetworks}
	VpcIpSubnetworks *[]*string `field:"optional" json:"vpcIpSubnetworks" yaml:"vpcIpSubnetworks"`
}

