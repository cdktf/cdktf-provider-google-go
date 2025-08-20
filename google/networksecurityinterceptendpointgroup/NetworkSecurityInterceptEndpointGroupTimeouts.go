// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityinterceptendpointgroup


type NetworkSecurityInterceptEndpointGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/network_security_intercept_endpoint_group#create NetworkSecurityInterceptEndpointGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/network_security_intercept_endpoint_group#delete NetworkSecurityInterceptEndpointGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/network_security_intercept_endpoint_group#update NetworkSecurityInterceptEndpointGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

