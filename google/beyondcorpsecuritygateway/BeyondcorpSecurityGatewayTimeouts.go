// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package beyondcorpsecuritygateway


type BeyondcorpSecurityGatewayTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/beyondcorp_security_gateway#create BeyondcorpSecurityGateway#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/beyondcorp_security_gateway#delete BeyondcorpSecurityGateway#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/beyondcorp_security_gateway#update BeyondcorpSecurityGateway#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

