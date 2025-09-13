// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apikeyskey


type ApikeysKeyRestrictions struct {
	// android_key_restrictions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apikeys_key#android_key_restrictions ApikeysKey#android_key_restrictions}
	AndroidKeyRestrictions *ApikeysKeyRestrictionsAndroidKeyRestrictions `field:"optional" json:"androidKeyRestrictions" yaml:"androidKeyRestrictions"`
	// api_targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apikeys_key#api_targets ApikeysKey#api_targets}
	ApiTargets interface{} `field:"optional" json:"apiTargets" yaml:"apiTargets"`
	// browser_key_restrictions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apikeys_key#browser_key_restrictions ApikeysKey#browser_key_restrictions}
	BrowserKeyRestrictions *ApikeysKeyRestrictionsBrowserKeyRestrictions `field:"optional" json:"browserKeyRestrictions" yaml:"browserKeyRestrictions"`
	// ios_key_restrictions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apikeys_key#ios_key_restrictions ApikeysKey#ios_key_restrictions}
	IosKeyRestrictions *ApikeysKeyRestrictionsIosKeyRestrictions `field:"optional" json:"iosKeyRestrictions" yaml:"iosKeyRestrictions"`
	// server_key_restrictions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apikeys_key#server_key_restrictions ApikeysKey#server_key_restrictions}
	ServerKeyRestrictions *ApikeysKeyRestrictionsServerKeyRestrictions `field:"optional" json:"serverKeyRestrictions" yaml:"serverKeyRestrictions"`
}

