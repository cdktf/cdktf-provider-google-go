// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelarmorfloorsetting


type ModelArmorFloorsettingAiPlatformFloorSetting struct {
	// If true, log Model Armor filter results to Cloud Logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/model_armor_floorsetting#enable_cloud_logging ModelArmorFloorsetting#enable_cloud_logging}
	EnableCloudLogging interface{} `field:"optional" json:"enableCloudLogging" yaml:"enableCloudLogging"`
	// If true, Model Armor filters will be run in inspect and block mode.
	//
	// Requests that trip Model Armor filters will be blocked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/model_armor_floorsetting#inspect_and_block ModelArmorFloorsetting#inspect_and_block}
	InspectAndBlock interface{} `field:"optional" json:"inspectAndBlock" yaml:"inspectAndBlock"`
	// If true, Model Armor filters will be run in inspect only mode. No action will be taken on the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/model_armor_floorsetting#inspect_only ModelArmorFloorsetting#inspect_only}
	InspectOnly interface{} `field:"optional" json:"inspectOnly" yaml:"inspectOnly"`
}

