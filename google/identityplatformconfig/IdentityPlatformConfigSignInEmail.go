// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatformconfig


type IdentityPlatformConfigSignInEmail struct {
	// Whether email auth is enabled for the project or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.12.0/docs/resources/identity_platform_config#enabled IdentityPlatformConfig#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Whether a password is required for email auth or not.
	//
	// If true, both an email and
	// password must be provided to sign in. If false, a user may sign in via either
	// email/password or email link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.12.0/docs/resources/identity_platform_config#password_required IdentityPlatformConfig#password_required}
	PasswordRequired interface{} `field:"optional" json:"passwordRequired" yaml:"passwordRequired"`
}

