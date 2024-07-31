// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computerouterpeer


type ComputeRouterPeerCustomLearnedIpRanges struct {
	// The IP range to advertise. The value must be a CIDR-formatted string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.1/docs/resources/compute_router_peer#range ComputeRouterPeer#range}
	Range *string `field:"required" json:"range" yaml:"range"`
}

