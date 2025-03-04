// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package binaryauthorizationattestor


type BinaryAuthorizationAttestorTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/binary_authorization_attestor#create BinaryAuthorizationAttestor#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/binary_authorization_attestor#delete BinaryAuthorizationAttestor#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/binary_authorization_attestor#update BinaryAuthorizationAttestor#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

