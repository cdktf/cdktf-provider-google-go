// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatformconfig


type IdentityPlatformConfigSignIn struct {
	// Whether to allow more than one account to have the same email.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.42.0/docs/resources/identity_platform_config#allow_duplicate_emails IdentityPlatformConfig#allow_duplicate_emails}
	AllowDuplicateEmails interface{} `field:"optional" json:"allowDuplicateEmails" yaml:"allowDuplicateEmails"`
	// anonymous block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.42.0/docs/resources/identity_platform_config#anonymous IdentityPlatformConfig#anonymous}
	Anonymous *IdentityPlatformConfigSignInAnonymous `field:"optional" json:"anonymous" yaml:"anonymous"`
	// email block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.42.0/docs/resources/identity_platform_config#email IdentityPlatformConfig#email}
	Email *IdentityPlatformConfigSignInEmail `field:"optional" json:"email" yaml:"email"`
	// phone_number block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.42.0/docs/resources/identity_platform_config#phone_number IdentityPlatformConfig#phone_number}
	PhoneNumber *IdentityPlatformConfigSignInPhoneNumber `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
}

