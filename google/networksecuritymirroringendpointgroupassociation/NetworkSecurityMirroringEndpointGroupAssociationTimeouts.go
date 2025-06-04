// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecuritymirroringendpointgroupassociation


type NetworkSecurityMirroringEndpointGroupAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/network_security_mirroring_endpoint_group_association#create NetworkSecurityMirroringEndpointGroupAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/network_security_mirroring_endpoint_group_association#delete NetworkSecurityMirroringEndpointGroupAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/network_security_mirroring_endpoint_group_association#update NetworkSecurityMirroringEndpointGroupAssociation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

