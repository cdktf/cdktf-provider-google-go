// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkconnectivitypolicybasedroute


type NetworkConnectivityPolicyBasedRouteVirtualMachine struct {
	// A list of VM instance tags that this policy-based route applies to.
	//
	// VM instances that have ANY of tags specified here will install this PBR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/network_connectivity_policy_based_route#tags NetworkConnectivityPolicyBasedRoute#tags}
	Tags *[]*string `field:"required" json:"tags" yaml:"tags"`
}

