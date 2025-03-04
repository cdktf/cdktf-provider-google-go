// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computerouterinterface


type ComputeRouterInterfaceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/compute_router_interface#create ComputeRouterInterface#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/compute_router_interface#delete ComputeRouterInterface#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

