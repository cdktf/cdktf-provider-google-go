// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package developerconnectconnection


type DeveloperConnectConnectionGitlabEnterpriseConfig struct {
	// authorizer_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/developer_connect_connection#authorizer_credential DeveloperConnectConnection#authorizer_credential}
	AuthorizerCredential *DeveloperConnectConnectionGitlabEnterpriseConfigAuthorizerCredential `field:"required" json:"authorizerCredential" yaml:"authorizerCredential"`
	// Required. The URI of the GitLab Enterprise host this connection is for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/developer_connect_connection#host_uri DeveloperConnectConnection#host_uri}
	HostUri *string `field:"required" json:"hostUri" yaml:"hostUri"`
	// read_authorizer_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/developer_connect_connection#read_authorizer_credential DeveloperConnectConnection#read_authorizer_credential}
	ReadAuthorizerCredential *DeveloperConnectConnectionGitlabEnterpriseConfigReadAuthorizerCredential `field:"required" json:"readAuthorizerCredential" yaml:"readAuthorizerCredential"`
	// Required.
	//
	// Immutable. SecretManager resource containing the webhook secret of a GitLab project,
	// formatted as 'projects/* /secrets/* /versions/*'. This is used to validate
	// webhooks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/developer_connect_connection#webhook_secret_secret_version DeveloperConnectConnection#webhook_secret_secret_version}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	WebhookSecretSecretVersion *string `field:"required" json:"webhookSecretSecretVersion" yaml:"webhookSecretSecretVersion"`
	// service_directory_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/developer_connect_connection#service_directory_config DeveloperConnectConnection#service_directory_config}
	ServiceDirectoryConfig *DeveloperConnectConnectionGitlabEnterpriseConfigServiceDirectoryConfig `field:"optional" json:"serviceDirectoryConfig" yaml:"serviceDirectoryConfig"`
	// Optional. SSL Certificate Authority certificate to use for requests to GitLab Enterprise instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/developer_connect_connection#ssl_ca_certificate DeveloperConnectConnection#ssl_ca_certificate}
	SslCaCertificate *string `field:"optional" json:"sslCaCertificate" yaml:"sslCaCertificate"`
}

