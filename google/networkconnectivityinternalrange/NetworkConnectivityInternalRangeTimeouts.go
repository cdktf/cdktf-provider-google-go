// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkconnectivityinternalrange


type NetworkConnectivityInternalRangeTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.35.0/docs/resources/network_connectivity_internal_range#create NetworkConnectivityInternalRange#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.35.0/docs/resources/network_connectivity_internal_range#delete NetworkConnectivityInternalRange#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.35.0/docs/resources/network_connectivity_internal_range#update NetworkConnectivityInternalRange#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

