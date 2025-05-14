// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityauthzpolicy


type NetworkSecurityAuthzPolicyHttpRulesFrom struct {
	// not_sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/network_security_authz_policy#not_sources NetworkSecurityAuthzPolicy#not_sources}
	NotSources interface{} `field:"optional" json:"notSources" yaml:"notSources"`
	// sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/network_security_authz_policy#sources NetworkSecurityAuthzPolicy#sources}
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
}

