// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelarmortemplate


type ModelArmorTemplateFilterConfig struct {
	// malicious_uri_filter_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/model_armor_template#malicious_uri_filter_settings ModelArmorTemplate#malicious_uri_filter_settings}
	MaliciousUriFilterSettings *ModelArmorTemplateFilterConfigMaliciousUriFilterSettings `field:"optional" json:"maliciousUriFilterSettings" yaml:"maliciousUriFilterSettings"`
	// pi_and_jailbreak_filter_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/model_armor_template#pi_and_jailbreak_filter_settings ModelArmorTemplate#pi_and_jailbreak_filter_settings}
	PiAndJailbreakFilterSettings *ModelArmorTemplateFilterConfigPiAndJailbreakFilterSettings `field:"optional" json:"piAndJailbreakFilterSettings" yaml:"piAndJailbreakFilterSettings"`
	// rai_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/model_armor_template#rai_settings ModelArmorTemplate#rai_settings}
	RaiSettings *ModelArmorTemplateFilterConfigRaiSettings `field:"optional" json:"raiSettings" yaml:"raiSettings"`
	// sdp_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/model_armor_template#sdp_settings ModelArmorTemplate#sdp_settings}
	SdpSettings *ModelArmorTemplateFilterConfigSdpSettings `field:"optional" json:"sdpSettings" yaml:"sdpSettings"`
}

