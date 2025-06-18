// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computerouter


type ComputeRouterMd5AuthenticationKeys struct {
	// Value of the key used for MD5 authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/compute_router#key ComputeRouter#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Name used to identify the key.
	//
	// Must be unique within a router.
	// Must be referenced by exactly one bgpPeer. Must comply with RFC1035.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/compute_router#name ComputeRouter#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

