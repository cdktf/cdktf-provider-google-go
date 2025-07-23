// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelarmorfloorsetting


type ModelArmorFloorsettingFilterConfigSdpSettings struct {
	// advanced_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/model_armor_floorsetting#advanced_config ModelArmorFloorsetting#advanced_config}
	AdvancedConfig *ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfig `field:"optional" json:"advancedConfig" yaml:"advancedConfig"`
	// basic_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/model_armor_floorsetting#basic_config ModelArmorFloorsetting#basic_config}
	BasicConfig *ModelArmorFloorsettingFilterConfigSdpSettingsBasicConfig `field:"optional" json:"basicConfig" yaml:"basicConfig"`
}

