// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkconnectivitypolicybasedroute


type NetworkConnectivityPolicyBasedRouteTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/network_connectivity_policy_based_route#create NetworkConnectivityPolicyBasedRoute#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/network_connectivity_policy_based_route#delete NetworkConnectivityPolicyBasedRoute#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

