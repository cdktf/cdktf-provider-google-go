// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatformconfig


type IdentityPlatformConfigMfaProviderConfigsTotpProviderConfig struct {
	// The allowed number of adjacent intervals that will be used for verification to avoid clock skew.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/identity_platform_config#adjacent_intervals IdentityPlatformConfig#adjacent_intervals}
	AdjacentIntervals *float64 `field:"optional" json:"adjacentIntervals" yaml:"adjacentIntervals"`
}

