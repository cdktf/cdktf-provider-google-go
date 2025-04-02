// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityauthzpolicy


type NetworkSecurityAuthzPolicyHttpRulesToOperations struct {
	// header_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/network_security_authz_policy#header_set NetworkSecurityAuthzPolicy#header_set}
	HeaderSet *NetworkSecurityAuthzPolicyHttpRulesToOperationsHeaderSet `field:"optional" json:"headerSet" yaml:"headerSet"`
	// hosts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/network_security_authz_policy#hosts NetworkSecurityAuthzPolicy#hosts}
	Hosts interface{} `field:"optional" json:"hosts" yaml:"hosts"`
	// A list of HTTP methods to match against.
	//
	// Each entry must be a valid HTTP method name (GET, PUT, POST, HEAD, PATCH, DELETE, OPTIONS). It only allows exact match and is always case sensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/network_security_authz_policy#methods NetworkSecurityAuthzPolicy#methods}
	Methods *[]*string `field:"optional" json:"methods" yaml:"methods"`
	// paths block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/network_security_authz_policy#paths NetworkSecurityAuthzPolicy#paths}
	Paths interface{} `field:"optional" json:"paths" yaml:"paths"`
}

