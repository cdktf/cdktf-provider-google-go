// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxwebhook


type DialogflowCxWebhookServiceDirectoryGenericWebServiceOauthConfig struct {
	// The client ID provided by the 3rd party platform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dialogflow_cx_webhook#client_id DialogflowCxWebhook#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The token endpoint provided by the 3rd party platform to exchange an access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dialogflow_cx_webhook#token_endpoint DialogflowCxWebhook#token_endpoint}
	TokenEndpoint *string `field:"required" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// The client secret provided by the 3rd party platform.  If the 'secret_version_for_client_secret' field is set, this field will be ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dialogflow_cx_webhook#client_secret DialogflowCxWebhook#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The OAuth scopes to grant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dialogflow_cx_webhook#scopes DialogflowCxWebhook#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	// The name of the SecretManager secret version resource storing the client secret.
	//
	// If this field is set, the 'client_secret' field will be
	// ignored.
	// Format: 'projects/{project}/secrets/{secret}/versions/{version}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dialogflow_cx_webhook#secret_version_for_client_secret DialogflowCxWebhook#secret_version_for_client_secret}
	SecretVersionForClientSecret *string `field:"optional" json:"secretVersionForClientSecret" yaml:"secretVersionForClientSecret"`
}

