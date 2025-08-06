// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxtool


type DialogflowCxToolOpenApiSpecAuthenticationApiKeyConfig struct {
	// The parameter name or the header name of the API key.
	//
	// E.g., If the API request is "https://example.com/act?X-Api-Key=", "X-Api-Key" would be the parameter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dialogflow_cx_tool#key_name DialogflowCxTool#key_name}
	KeyName *string `field:"required" json:"keyName" yaml:"keyName"`
	// Key location in the request. See [RequestLocation](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.tools#requestlocation) for valid values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dialogflow_cx_tool#request_location DialogflowCxTool#request_location}
	RequestLocation *string `field:"required" json:"requestLocation" yaml:"requestLocation"`
	// Optional. The API key. If the 'secretVersionForApiKey'' field is set, this field will be ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dialogflow_cx_tool#api_key DialogflowCxTool#api_key}
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Optional.
	//
	// The name of the SecretManager secret version resource storing the API key.
	// If this field is set, the apiKey field will be ignored.
	// Format: projects/{project}/secrets/{secret}/versions/{version}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dialogflow_cx_tool#secret_version_for_api_key DialogflowCxTool#secret_version_for_api_key}
	SecretVersionForApiKey *string `field:"optional" json:"secretVersionForApiKey" yaml:"secretVersionForApiKey"`
}

