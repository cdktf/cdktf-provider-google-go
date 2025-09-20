// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityauthzpolicy


type NetworkSecurityAuthzPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/network_security_authz_policy#create NetworkSecurityAuthzPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/network_security_authz_policy#delete NetworkSecurityAuthzPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/network_security_authz_policy#update NetworkSecurityAuthzPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

