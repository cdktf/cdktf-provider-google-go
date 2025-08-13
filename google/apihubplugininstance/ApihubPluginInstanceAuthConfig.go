// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apihubplugininstance


type ApihubPluginInstanceAuthConfig struct {
	// Possible values: AUTH_TYPE_UNSPECIFIED NO_AUTH GOOGLE_SERVICE_ACCOUNT USER_PASSWORD API_KEY OAUTH2_CLIENT_CREDENTIALS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apihub_plugin_instance#auth_type ApihubPluginInstance#auth_type}
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// api_key_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apihub_plugin_instance#api_key_config ApihubPluginInstance#api_key_config}
	ApiKeyConfig *ApihubPluginInstanceAuthConfigApiKeyConfig `field:"optional" json:"apiKeyConfig" yaml:"apiKeyConfig"`
	// google_service_account_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apihub_plugin_instance#google_service_account_config ApihubPluginInstance#google_service_account_config}
	GoogleServiceAccountConfig *ApihubPluginInstanceAuthConfigGoogleServiceAccountConfig `field:"optional" json:"googleServiceAccountConfig" yaml:"googleServiceAccountConfig"`
	// oauth2_client_credentials_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apihub_plugin_instance#oauth2_client_credentials_config ApihubPluginInstance#oauth2_client_credentials_config}
	Oauth2ClientCredentialsConfig *ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfig `field:"optional" json:"oauth2ClientCredentialsConfig" yaml:"oauth2ClientCredentialsConfig"`
	// user_password_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apihub_plugin_instance#user_password_config ApihubPluginInstance#user_password_config}
	UserPasswordConfig *ApihubPluginInstanceAuthConfigUserPasswordConfig `field:"optional" json:"userPasswordConfig" yaml:"userPasswordConfig"`
}

