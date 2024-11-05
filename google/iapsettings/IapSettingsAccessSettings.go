// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iapsettings


type IapSettingsAccessSettings struct {
	// allowed_domains_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/iap_settings#allowed_domains_settings IapSettings#allowed_domains_settings}
	AllowedDomainsSettings *IapSettingsAccessSettingsAllowedDomainsSettings `field:"optional" json:"allowedDomainsSettings" yaml:"allowedDomainsSettings"`
	// cors_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/iap_settings#cors_settings IapSettings#cors_settings}
	CorsSettings *IapSettingsAccessSettingsCorsSettings `field:"optional" json:"corsSettings" yaml:"corsSettings"`
	// gcip_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/iap_settings#gcip_settings IapSettings#gcip_settings}
	GcipSettings *IapSettingsAccessSettingsGcipSettings `field:"optional" json:"gcipSettings" yaml:"gcipSettings"`
	// Identity sources that IAP can use to authenticate the end user.
	//
	// Only one identity source
	// can be configured. The possible values are:
	//
	// * 'WORKFORCE_IDENTITY_FEDERATION': Use external identities set up on Google Cloud Workforce
	//   				     Identity Federation. Possible values: ["WORKFORCE_IDENTITY_FEDERATION"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/iap_settings#identity_sources IapSettings#identity_sources}
	IdentitySources *[]*string `field:"optional" json:"identitySources" yaml:"identitySources"`
	// oauth_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/iap_settings#oauth_settings IapSettings#oauth_settings}
	OauthSettings *IapSettingsAccessSettingsOauthSettings `field:"optional" json:"oauthSettings" yaml:"oauthSettings"`
	// reauth_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/iap_settings#reauth_settings IapSettings#reauth_settings}
	ReauthSettings *IapSettingsAccessSettingsReauthSettings `field:"optional" json:"reauthSettings" yaml:"reauthSettings"`
	// workforce_identity_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/iap_settings#workforce_identity_settings IapSettings#workforce_identity_settings}
	WorkforceIdentitySettings *IapSettingsAccessSettingsWorkforceIdentitySettings `field:"optional" json:"workforceIdentitySettings" yaml:"workforceIdentitySettings"`
}

