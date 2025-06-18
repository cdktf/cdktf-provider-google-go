// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatformconfig


type IdentityPlatformConfigClientPermissions struct {
	// When true, end users cannot delete their account on the associated project through any of our API methods.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/identity_platform_config#disabled_user_deletion IdentityPlatformConfig#disabled_user_deletion}
	DisabledUserDeletion interface{} `field:"optional" json:"disabledUserDeletion" yaml:"disabledUserDeletion"`
	// When true, end users cannot sign up for a new account on the associated project through any of our API methods.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/identity_platform_config#disabled_user_signup IdentityPlatformConfig#disabled_user_signup}
	DisabledUserSignup interface{} `field:"optional" json:"disabledUserSignup" yaml:"disabledUserSignup"`
}

