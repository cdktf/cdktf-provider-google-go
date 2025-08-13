// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iapsettings


type IapSettingsApplicationSettingsAccessDeniedPageSettings struct {
	// The URI to be redirected to when access is denied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/iap_settings#access_denied_page_uri IapSettings#access_denied_page_uri}
	AccessDeniedPageUri *string `field:"optional" json:"accessDeniedPageUri" yaml:"accessDeniedPageUri"`
	// Whether to generate a troubleshooting URL on access denied events to this application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/iap_settings#generate_troubleshooting_uri IapSettings#generate_troubleshooting_uri}
	GenerateTroubleshootingUri interface{} `field:"optional" json:"generateTroubleshootingUri" yaml:"generateTroubleshootingUri"`
	// Whether to generate remediation token on access denied events to this application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/iap_settings#remediation_token_generation_enabled IapSettings#remediation_token_generation_enabled}
	RemediationTokenGenerationEnabled interface{} `field:"optional" json:"remediationTokenGenerationEnabled" yaml:"remediationTokenGenerationEnabled"`
}

