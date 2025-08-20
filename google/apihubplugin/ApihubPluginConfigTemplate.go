// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apihubplugin


type ApihubPluginConfigTemplate struct {
	// additional_config_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/apihub_plugin#additional_config_template ApihubPlugin#additional_config_template}
	AdditionalConfigTemplate interface{} `field:"optional" json:"additionalConfigTemplate" yaml:"additionalConfigTemplate"`
	// auth_config_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/apihub_plugin#auth_config_template ApihubPlugin#auth_config_template}
	AuthConfigTemplate *ApihubPluginConfigTemplateAuthConfigTemplate `field:"optional" json:"authConfigTemplate" yaml:"authConfigTemplate"`
}

