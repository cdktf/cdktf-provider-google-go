// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apihubplugininstance


type ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfig struct {
	// The client identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/apihub_plugin_instance#client_id ApihubPluginInstance#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// client_secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/apihub_plugin_instance#client_secret ApihubPluginInstance#client_secret}
	ClientSecret *ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecret `field:"required" json:"clientSecret" yaml:"clientSecret"`
}

