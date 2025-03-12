// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityservertlspolicy


type NetworkSecurityServerTlsPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/network_security_server_tls_policy#create NetworkSecurityServerTlsPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/network_security_server_tls_policy#delete NetworkSecurityServerTlsPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/network_security_server_tls_policy#update NetworkSecurityServerTlsPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

