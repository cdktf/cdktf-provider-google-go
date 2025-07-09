// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apihubplugin


type ApihubPluginConfigTemplateAuthConfigTemplate struct {
	// The list of authentication types supported by the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/apihub_plugin#supported_auth_types ApihubPlugin#supported_auth_types}
	SupportedAuthTypes *[]*string `field:"required" json:"supportedAuthTypes" yaml:"supportedAuthTypes"`
	// service_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/apihub_plugin#service_account ApihubPlugin#service_account}
	ServiceAccount *ApihubPluginConfigTemplateAuthConfigTemplateServiceAccount `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
}

