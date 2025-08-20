// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelarmorfloorsetting


type ModelArmorFloorsettingFilterConfigPiAndJailbreakFilterSettings struct {
	// Possible values: LOW_AND_ABOVE MEDIUM_AND_ABOVE HIGH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/model_armor_floorsetting#confidence_level ModelArmorFloorsetting#confidence_level}
	ConfidenceLevel *string `field:"optional" json:"confidenceLevel" yaml:"confidenceLevel"`
	// Tells whether Prompt injection and Jailbreak filter is enabled or disabled. Possible values: ENABLED DISABLED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/model_armor_floorsetting#filter_enforcement ModelArmorFloorsetting#filter_enforcement}
	FilterEnforcement *string `field:"optional" json:"filterEnforcement" yaml:"filterEnforcement"`
}

