// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatformconfig


type IdentityPlatformConfigMfaProviderConfigs struct {
	// Whether MultiFactor Authentication has been enabled for this project. Possible values: ["DISABLED", "ENABLED", "MANDATORY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/identity_platform_config#state IdentityPlatformConfig#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// totp_provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/identity_platform_config#totp_provider_config IdentityPlatformConfig#totp_provider_config}
	TotpProviderConfig *IdentityPlatformConfigMfaProviderConfigsTotpProviderConfig `field:"optional" json:"totpProviderConfig" yaml:"totpProviderConfig"`
}

