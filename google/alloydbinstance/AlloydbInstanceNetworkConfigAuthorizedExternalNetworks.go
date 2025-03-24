// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alloydbinstance


type AlloydbInstanceNetworkConfigAuthorizedExternalNetworks struct {
	// CIDR range for one authorized network of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/alloydb_instance#cidr_range AlloydbInstance#cidr_range}
	CidrRange *string `field:"optional" json:"cidrRange" yaml:"cidrRange"`
}

