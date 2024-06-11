// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatformconfig


type IdentityPlatformConfigBlockingFunctions struct {
	// triggers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.33.0/docs/resources/identity_platform_config#triggers IdentityPlatformConfig#triggers}
	Triggers interface{} `field:"required" json:"triggers" yaml:"triggers"`
	// forward_inbound_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.33.0/docs/resources/identity_platform_config#forward_inbound_credentials IdentityPlatformConfig#forward_inbound_credentials}
	ForwardInboundCredentials *IdentityPlatformConfigBlockingFunctionsForwardInboundCredentials `field:"optional" json:"forwardInboundCredentials" yaml:"forwardInboundCredentials"`
}

