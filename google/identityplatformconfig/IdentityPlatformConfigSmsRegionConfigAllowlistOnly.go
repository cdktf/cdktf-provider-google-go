// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatformconfig


type IdentityPlatformConfigSmsRegionConfigAllowlistOnly struct {
	// Two letter unicode region codes to allow as defined by https://cldr.unicode.org/ The full list of these region codes is here: https://github.com/unicode-cldr/cldr-localenames-full/blob/master/main/en/territories.json.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/identity_platform_config#allowed_regions IdentityPlatformConfig#allowed_regions}
	AllowedRegions *[]*string `field:"optional" json:"allowedRegions" yaml:"allowedRegions"`
}

