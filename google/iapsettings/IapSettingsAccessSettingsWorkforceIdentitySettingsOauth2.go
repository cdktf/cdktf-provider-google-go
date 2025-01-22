// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iapsettings


type IapSettingsAccessSettingsWorkforceIdentitySettingsOauth2 struct {
	// The OAuth 2.0 client ID registered in the workforce identity federation OAuth 2.0 Server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/iap_settings#client_id IapSettings#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Input only. The OAuth 2.0 client secret created while registering the client ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/iap_settings#client_secret IapSettings#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
}

