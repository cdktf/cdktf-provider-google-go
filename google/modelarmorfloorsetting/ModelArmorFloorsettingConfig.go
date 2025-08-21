// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelarmorfloorsetting

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ModelArmorFloorsettingConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// filter_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/model_armor_floorsetting#filter_config ModelArmorFloorsetting#filter_config}
	FilterConfig *ModelArmorFloorsettingFilterConfig `field:"required" json:"filterConfig" yaml:"filterConfig"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/model_armor_floorsetting#location ModelArmorFloorsetting#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Will be any one of these:.
	//
	// * 'projects/{project}'
	// * 'folders/{folder}'
	// * 'organizations/{organizationId}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/model_armor_floorsetting#parent ModelArmorFloorsetting#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// ai_platform_floor_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/model_armor_floorsetting#ai_platform_floor_setting ModelArmorFloorsetting#ai_platform_floor_setting}
	AiPlatformFloorSetting *ModelArmorFloorsettingAiPlatformFloorSetting `field:"optional" json:"aiPlatformFloorSetting" yaml:"aiPlatformFloorSetting"`
	// Floor Settings enforcement status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/model_armor_floorsetting#enable_floor_setting_enforcement ModelArmorFloorsetting#enable_floor_setting_enforcement}
	EnableFloorSettingEnforcement interface{} `field:"optional" json:"enableFloorSettingEnforcement" yaml:"enableFloorSettingEnforcement"`
	// floor_setting_metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/model_armor_floorsetting#floor_setting_metadata ModelArmorFloorsetting#floor_setting_metadata}
	FloorSettingMetadata *ModelArmorFloorsettingFloorSettingMetadata `field:"optional" json:"floorSettingMetadata" yaml:"floorSettingMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/model_armor_floorsetting#id ModelArmorFloorsetting#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// List of integrated services for which the floor setting is applicable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/model_armor_floorsetting#integrated_services ModelArmorFloorsetting#integrated_services}
	IntegratedServices *[]*string `field:"optional" json:"integratedServices" yaml:"integratedServices"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/model_armor_floorsetting#timeouts ModelArmorFloorsetting#timeouts}
	Timeouts *ModelArmorFloorsettingTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

