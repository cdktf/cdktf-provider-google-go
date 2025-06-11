// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computerouternat


type ComputeRouterNatNat64Subnetwork struct {
	// Self-link of the subnetwork resource that will use NAT64.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/compute_router_nat#name ComputeRouterNat#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

