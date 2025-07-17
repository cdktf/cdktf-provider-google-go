// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iapsettings


type IapSettingsAccessSettingsCorsSettings struct {
	// Configuration to allow HTTP OPTIONS calls to skip authorization.
	//
	// If undefined, IAP will not apply any special logic to OPTIONS requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/iap_settings#allow_http_options IapSettings#allow_http_options}
	AllowHttpOptions interface{} `field:"optional" json:"allowHttpOptions" yaml:"allowHttpOptions"`
}

