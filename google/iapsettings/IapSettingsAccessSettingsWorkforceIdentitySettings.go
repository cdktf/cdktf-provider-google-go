// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iapsettings


type IapSettingsAccessSettingsWorkforceIdentitySettings struct {
	// oauth2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.16.0/docs/resources/iap_settings#oauth2 IapSettings#oauth2}
	Oauth2 *IapSettingsAccessSettingsWorkforceIdentitySettingsOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
	// The workforce pool resources. Only one workforce pool is accepted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.16.0/docs/resources/iap_settings#workforce_pools IapSettings#workforce_pools}
	WorkforcePools *[]*string `field:"optional" json:"workforcePools" yaml:"workforcePools"`
}

