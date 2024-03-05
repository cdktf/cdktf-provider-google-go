// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatformconfig


type IdentityPlatformConfigSmsRegionConfigAllowByDefault struct {
	// Two letter unicode region codes to disallow as defined by https://cldr.unicode.org/ The full list of these region codes is here: https://github.com/unicode-cldr/cldr-localenames-full/blob/master/main/en/territories.json.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/identity_platform_config#disallowed_regions IdentityPlatformConfig#disallowed_regions}
	DisallowedRegions *[]*string `field:"optional" json:"disallowedRegions" yaml:"disallowedRegions"`
}

