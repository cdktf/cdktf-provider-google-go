// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityauthzpolicy


type NetworkSecurityAuthzPolicyHttpRulesToOperationsHeaderSetHeaders struct {
	// Specifies the name of the header in the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/network_security_authz_policy#name NetworkSecurityAuthzPolicy#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/network_security_authz_policy#value NetworkSecurityAuthzPolicy#value}
	Value *NetworkSecurityAuthzPolicyHttpRulesToOperationsHeaderSetHeadersValue `field:"optional" json:"value" yaml:"value"`
}

