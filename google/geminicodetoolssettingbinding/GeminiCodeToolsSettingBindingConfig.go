// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package geminicodetoolssettingbinding

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GeminiCodeToolsSettingBindingConfig struct {
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
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/gemini_code_tools_setting_binding#code_tools_setting_id GeminiCodeToolsSettingBinding#code_tools_setting_id}
	CodeToolsSettingId *string `field:"required" json:"codeToolsSettingId" yaml:"codeToolsSettingId"`
	// Id of the setting binding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/gemini_code_tools_setting_binding#setting_binding_id GeminiCodeToolsSettingBinding#setting_binding_id}
	SettingBindingId *string `field:"required" json:"settingBindingId" yaml:"settingBindingId"`
	// Target of the binding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/gemini_code_tools_setting_binding#target GeminiCodeToolsSettingBinding#target}
	Target *string `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/gemini_code_tools_setting_binding#id GeminiCodeToolsSettingBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/gemini_code_tools_setting_binding#labels GeminiCodeToolsSettingBinding#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/gemini_code_tools_setting_binding#location GeminiCodeToolsSettingBinding#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Product type of the setting binding. Possible values: ["GEMINI_CODE_ASSIST"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/gemini_code_tools_setting_binding#product GeminiCodeToolsSettingBinding#product}
	Product *string `field:"optional" json:"product" yaml:"product"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/gemini_code_tools_setting_binding#project GeminiCodeToolsSettingBinding#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/gemini_code_tools_setting_binding#timeouts GeminiCodeToolsSettingBinding#timeouts}
	Timeouts *GeminiCodeToolsSettingBindingTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

