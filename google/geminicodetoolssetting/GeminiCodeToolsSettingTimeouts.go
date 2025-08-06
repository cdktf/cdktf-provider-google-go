// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package geminicodetoolssetting


type GeminiCodeToolsSettingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/gemini_code_tools_setting#create GeminiCodeToolsSetting#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/gemini_code_tools_setting#delete GeminiCodeToolsSetting#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/gemini_code_tools_setting#update GeminiCodeToolsSetting#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

