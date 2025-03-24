// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityclienttlspolicy


type NetworkSecurityClientTlsPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/network_security_client_tls_policy#create NetworkSecurityClientTlsPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/network_security_client_tls_policy#delete NetworkSecurityClientTlsPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/network_security_client_tls_policy#update NetworkSecurityClientTlsPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

