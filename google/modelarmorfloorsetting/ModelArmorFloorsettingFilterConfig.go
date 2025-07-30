// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelarmorfloorsetting


type ModelArmorFloorsettingFilterConfig struct {
	// malicious_uri_filter_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/model_armor_floorsetting#malicious_uri_filter_settings ModelArmorFloorsetting#malicious_uri_filter_settings}
	MaliciousUriFilterSettings *ModelArmorFloorsettingFilterConfigMaliciousUriFilterSettings `field:"optional" json:"maliciousUriFilterSettings" yaml:"maliciousUriFilterSettings"`
	// pi_and_jailbreak_filter_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/model_armor_floorsetting#pi_and_jailbreak_filter_settings ModelArmorFloorsetting#pi_and_jailbreak_filter_settings}
	PiAndJailbreakFilterSettings *ModelArmorFloorsettingFilterConfigPiAndJailbreakFilterSettings `field:"optional" json:"piAndJailbreakFilterSettings" yaml:"piAndJailbreakFilterSettings"`
	// rai_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/model_armor_floorsetting#rai_settings ModelArmorFloorsetting#rai_settings}
	RaiSettings *ModelArmorFloorsettingFilterConfigRaiSettings `field:"optional" json:"raiSettings" yaml:"raiSettings"`
	// sdp_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/model_armor_floorsetting#sdp_settings ModelArmorFloorsetting#sdp_settings}
	SdpSettings *ModelArmorFloorsettingFilterConfigSdpSettings `field:"optional" json:"sdpSettings" yaml:"sdpSettings"`
}

