// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecuritymirroringendpointgroup


type NetworkSecurityMirroringEndpointGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/network_security_mirroring_endpoint_group#create NetworkSecurityMirroringEndpointGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/network_security_mirroring_endpoint_group#delete NetworkSecurityMirroringEndpointGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/network_security_mirroring_endpoint_group#update NetworkSecurityMirroringEndpointGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

