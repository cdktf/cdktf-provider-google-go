// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecuritybackendauthenticationconfig


type NetworkSecurityBackendAuthenticationConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/network_security_backend_authentication_config#create NetworkSecurityBackendAuthenticationConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/network_security_backend_authentication_config#delete NetworkSecurityBackendAuthenticationConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/network_security_backend_authentication_config#update NetworkSecurityBackendAuthenticationConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

