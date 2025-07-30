// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxwebhook


type DialogflowCxWebhookServiceDirectoryGenericWebService struct {
	// The webhook URI for receiving POST requests. It must use https protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_webhook#uri DialogflowCxWebhook#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// Specifies a list of allowed custom CA certificates (in DER format) for HTTPS verification.
	//
	// This overrides the default SSL trust store. If this
	// is empty or unspecified, Dialogflow will use Google's default trust store
	// to verify certificates.
	// N.B. Make sure the HTTPS server certificates are signed with "subject alt
	// name". For instance a certificate can be self-signed using the following
	// command,
	// openssl x509 -req -days 200 -in example.com.csr \
	// -signkey example.com.key \
	// -out example.com.crt \
	// -extfile <(printf "\nsubjectAltName='DNS:www.example.com'")
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_webhook#allowed_ca_certs DialogflowCxWebhook#allowed_ca_certs}
	AllowedCaCerts *[]*string `field:"optional" json:"allowedCaCerts" yaml:"allowedCaCerts"`
	// HTTP method for the flexible webhook calls.
	//
	// Standard webhook always uses
	// POST. Possible values: ["POST", "GET", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_webhook#http_method DialogflowCxWebhook#http_method}
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// oauth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_webhook#oauth_config DialogflowCxWebhook#oauth_config}
	OauthConfig *DialogflowCxWebhookServiceDirectoryGenericWebServiceOauthConfig `field:"optional" json:"oauthConfig" yaml:"oauthConfig"`
	// Maps the values extracted from specific fields of the flexible webhook response into session parameters.
	//
	// - Key: session parameter name
	// - Value: field path in the webhook response
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_webhook#parameter_mapping DialogflowCxWebhook#parameter_mapping}
	ParameterMapping *map[string]*string `field:"optional" json:"parameterMapping" yaml:"parameterMapping"`
	// Defines a custom JSON object as request body to send to flexible webhook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_webhook#request_body DialogflowCxWebhook#request_body}
	RequestBody *string `field:"optional" json:"requestBody" yaml:"requestBody"`
	// The HTTP request headers to send together with webhook requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_webhook#request_headers DialogflowCxWebhook#request_headers}
	RequestHeaders *map[string]*string `field:"optional" json:"requestHeaders" yaml:"requestHeaders"`
	// The SecretManager secret version resource storing the username:password pair for HTTP Basic authentication. Format: 'projects/{project}/secrets/{secret}/versions/{version}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_webhook#secret_version_for_username_password DialogflowCxWebhook#secret_version_for_username_password}
	SecretVersionForUsernamePassword *string `field:"optional" json:"secretVersionForUsernamePassword" yaml:"secretVersionForUsernamePassword"`
	// secret_versions_for_request_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_webhook#secret_versions_for_request_headers DialogflowCxWebhook#secret_versions_for_request_headers}
	SecretVersionsForRequestHeaders interface{} `field:"optional" json:"secretVersionsForRequestHeaders" yaml:"secretVersionsForRequestHeaders"`
	// Indicate the auth token type generated from the [Diglogflow service agent](https://cloud.google.com/iam/docs/service-agents#dialogflow-service-agent). The generated token is sent in the Authorization header. Possible values: ["NONE", "ID_TOKEN", "ACCESS_TOKEN"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_webhook#service_agent_auth DialogflowCxWebhook#service_agent_auth}
	ServiceAgentAuth *string `field:"optional" json:"serviceAgentAuth" yaml:"serviceAgentAuth"`
	// Type of the webhook. Possible values: ["STANDARD", "FLEXIBLE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_webhook#webhook_type DialogflowCxWebhook#webhook_type}
	WebhookType *string `field:"optional" json:"webhookType" yaml:"webhookType"`
}

