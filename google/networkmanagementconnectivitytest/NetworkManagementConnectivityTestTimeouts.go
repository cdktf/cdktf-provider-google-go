// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagementconnectivitytest


type NetworkManagementConnectivityTestTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/network_management_connectivity_test#create NetworkManagementConnectivityTest#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/network_management_connectivity_test#delete NetworkManagementConnectivityTest#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/network_management_connectivity_test#update NetworkManagementConnectivityTest#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

