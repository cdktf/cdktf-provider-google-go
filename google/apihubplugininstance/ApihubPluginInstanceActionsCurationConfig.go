// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apihubplugininstance


type ApihubPluginInstanceActionsCurationConfig struct {
	// Possible values: CURATION_TYPE_UNSPECIFIED DEFAULT_CURATION_FOR_API_METADATA CUSTOM_CURATION_FOR_API_METADATA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/apihub_plugin_instance#curation_type ApihubPluginInstance#curation_type}
	CurationType *string `field:"optional" json:"curationType" yaml:"curationType"`
	// custom_curation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/apihub_plugin_instance#custom_curation ApihubPluginInstance#custom_curation}
	CustomCuration *ApihubPluginInstanceActionsCurationConfigCustomCuration `field:"optional" json:"customCuration" yaml:"customCuration"`
}

