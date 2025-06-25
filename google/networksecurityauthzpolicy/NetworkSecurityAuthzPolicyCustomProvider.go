// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityauthzpolicy


type NetworkSecurityAuthzPolicyCustomProvider struct {
	// authz_extension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/network_security_authz_policy#authz_extension NetworkSecurityAuthzPolicy#authz_extension}
	AuthzExtension *NetworkSecurityAuthzPolicyCustomProviderAuthzExtension `field:"optional" json:"authzExtension" yaml:"authzExtension"`
	// cloud_iap block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/network_security_authz_policy#cloud_iap NetworkSecurityAuthzPolicy#cloud_iap}
	CloudIap *NetworkSecurityAuthzPolicyCustomProviderCloudIap `field:"optional" json:"cloudIap" yaml:"cloudIap"`
}

