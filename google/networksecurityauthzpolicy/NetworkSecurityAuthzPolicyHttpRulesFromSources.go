// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityauthzpolicy


type NetworkSecurityAuthzPolicyHttpRulesFromSources struct {
	// principals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/network_security_authz_policy#principals NetworkSecurityAuthzPolicy#principals}
	Principals interface{} `field:"optional" json:"principals" yaml:"principals"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/network_security_authz_policy#resources NetworkSecurityAuthzPolicy#resources}
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
}

