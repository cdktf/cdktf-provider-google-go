// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatformconfig


type IdentityPlatformConfigSignInPhoneNumber struct {
	// Whether phone number auth is enabled for the project or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/identity_platform_config#enabled IdentityPlatformConfig#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// A map of <test phone number, fake code> that can be used for phone auth testing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/identity_platform_config#test_phone_numbers IdentityPlatformConfig#test_phone_numbers}
	TestPhoneNumbers *map[string]*string `field:"optional" json:"testPhoneNumbers" yaml:"testPhoneNumbers"`
}

