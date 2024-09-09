// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatformconfig


type IdentityPlatformConfigMfa struct {
	// A list of usable second factors for this project. Possible values: ["PHONE_SMS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/identity_platform_config#enabled_providers IdentityPlatformConfig#enabled_providers}
	EnabledProviders *[]*string `field:"optional" json:"enabledProviders" yaml:"enabledProviders"`
	// provider_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/identity_platform_config#provider_configs IdentityPlatformConfig#provider_configs}
	ProviderConfigs interface{} `field:"optional" json:"providerConfigs" yaml:"providerConfigs"`
	// Whether MultiFactor Authentication has been enabled for this project. Possible values: ["DISABLED", "ENABLED", "MANDATORY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/identity_platform_config#state IdentityPlatformConfig#state}
	State *string `field:"optional" json:"state" yaml:"state"`
}

