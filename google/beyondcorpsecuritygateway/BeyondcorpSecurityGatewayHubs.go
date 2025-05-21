// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package beyondcorpsecuritygateway


type BeyondcorpSecurityGatewayHubs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/beyondcorp_security_gateway#region BeyondcorpSecurityGateway#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// internet_gateway block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/beyondcorp_security_gateway#internet_gateway BeyondcorpSecurityGateway#internet_gateway}
	InternetGateway *BeyondcorpSecurityGatewayHubsInternetGateway `field:"optional" json:"internetGateway" yaml:"internetGateway"`
}

