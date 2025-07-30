// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package beyondcorpsecuritygatewayapplication


type BeyondcorpSecurityGatewayApplicationUpstreamsEgressPolicy struct {
	// Required. List of regions where the application sends traffic to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/beyondcorp_security_gateway_application#regions BeyondcorpSecurityGatewayApplication#regions}
	Regions *[]*string `field:"required" json:"regions" yaml:"regions"`
}

