// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computerouterpeer


type ComputeRouterPeerMd5AuthenticationKey struct {
	// Value of the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/compute_router_peer#key ComputeRouterPeer#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// [REQUIRED] Name used to identify the key.
	//
	// Must be unique within a router. Must be referenced by exactly one bgpPeer. Must comply with RFC1035.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/compute_router_peer#name ComputeRouterPeer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

