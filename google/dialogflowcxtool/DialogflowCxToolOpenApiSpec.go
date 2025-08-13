// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxtool


type DialogflowCxToolOpenApiSpec struct {
	// The OpenAPI schema specified as a text.
	//
	// This field is part of a union field 'schema': only one of 'textSchema' may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dialogflow_cx_tool#text_schema DialogflowCxTool#text_schema}
	TextSchema *string `field:"required" json:"textSchema" yaml:"textSchema"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dialogflow_cx_tool#authentication DialogflowCxTool#authentication}
	Authentication *DialogflowCxToolOpenApiSpecAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// service_directory_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dialogflow_cx_tool#service_directory_config DialogflowCxTool#service_directory_config}
	ServiceDirectoryConfig *DialogflowCxToolOpenApiSpecServiceDirectoryConfig `field:"optional" json:"serviceDirectoryConfig" yaml:"serviceDirectoryConfig"`
	// tls_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dialogflow_cx_tool#tls_config DialogflowCxTool#tls_config}
	TlsConfig *DialogflowCxToolOpenApiSpecTlsConfig `field:"optional" json:"tlsConfig" yaml:"tlsConfig"`
}

