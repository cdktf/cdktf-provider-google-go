// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkconnectivitypolicybasedroute


type NetworkConnectivityPolicyBasedRouteInterconnectAttachment struct {
	// Cloud region to install this policy-based route on for Interconnect attachments.
	//
	// Use 'all' to install it on all Interconnect attachments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/network_connectivity_policy_based_route#region NetworkConnectivityPolicyBasedRoute#region}
	Region *string `field:"required" json:"region" yaml:"region"`
}

