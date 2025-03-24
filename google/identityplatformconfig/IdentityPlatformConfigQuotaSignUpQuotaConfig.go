// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatformconfig


type IdentityPlatformConfigQuotaSignUpQuotaConfig struct {
	// A sign up APIs quota that customers can override temporarily. Value can be in between 1 and 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/identity_platform_config#quota IdentityPlatformConfig#quota}
	Quota *float64 `field:"optional" json:"quota" yaml:"quota"`
	// How long this quota will be active for. It is measurred in seconds, e.g., Example: "9.615s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/identity_platform_config#quota_duration IdentityPlatformConfig#quota_duration}
	QuotaDuration *string `field:"optional" json:"quotaDuration" yaml:"quotaDuration"`
	// When this quota will take affect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/identity_platform_config#start_time IdentityPlatformConfig#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

