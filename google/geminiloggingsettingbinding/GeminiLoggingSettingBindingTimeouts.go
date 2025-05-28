// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package geminiloggingsettingbinding


type GeminiLoggingSettingBindingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/gemini_logging_setting_binding#create GeminiLoggingSettingBinding#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/gemini_logging_setting_binding#delete GeminiLoggingSettingBinding#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/gemini_logging_setting_binding#update GeminiLoggingSettingBinding#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

