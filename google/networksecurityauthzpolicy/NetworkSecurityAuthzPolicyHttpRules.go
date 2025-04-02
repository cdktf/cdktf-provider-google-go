// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityauthzpolicy


type NetworkSecurityAuthzPolicyHttpRules struct {
	// from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/network_security_authz_policy#from NetworkSecurityAuthzPolicy#from}
	From *NetworkSecurityAuthzPolicyHttpRulesFrom `field:"optional" json:"from" yaml:"from"`
	// to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/network_security_authz_policy#to NetworkSecurityAuthzPolicy#to}
	To *NetworkSecurityAuthzPolicyHttpRulesTo `field:"optional" json:"to" yaml:"to"`
	// CEL expression that describes the conditions to be satisfied for the action.
	//
	// The result of the CEL expression is ANDed with the from and to. Refer to the CEL language reference for a list of available attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/network_security_authz_policy#when NetworkSecurityAuthzPolicy#when}
	When *string `field:"optional" json:"when" yaml:"when"`
}

