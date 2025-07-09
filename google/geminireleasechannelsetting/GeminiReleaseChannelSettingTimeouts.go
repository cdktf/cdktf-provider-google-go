// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package geminireleasechannelsetting


type GeminiReleaseChannelSettingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/gemini_release_channel_setting#create GeminiReleaseChannelSetting#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/gemini_release_channel_setting#delete GeminiReleaseChannelSetting#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/gemini_release_channel_setting#update GeminiReleaseChannelSetting#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

