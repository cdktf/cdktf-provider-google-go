// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatformprojectdefaultconfig


type IdentityPlatformProjectDefaultConfigSignInAnonymous struct {
	// Whether anonymous user auth is enabled for the project or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.26.0/docs/resources/identity_platform_project_default_config#enabled IdentityPlatformProjectDefaultConfig#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

