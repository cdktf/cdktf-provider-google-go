// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package developerconnectconnection


type DeveloperConnectConnectionGithubEnterpriseConfig struct {
	// Required. The URI of the GitHub Enterprise host this connection is for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/developer_connect_connection#host_uri DeveloperConnectConnection#host_uri}
	HostUri *string `field:"required" json:"hostUri" yaml:"hostUri"`
	// Optional. ID of the GitHub App created from the manifest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/developer_connect_connection#app_id DeveloperConnectConnection#app_id}
	AppId *string `field:"optional" json:"appId" yaml:"appId"`
	// Optional. ID of the installation of the GitHub App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/developer_connect_connection#app_installation_id DeveloperConnectConnection#app_installation_id}
	AppInstallationId *string `field:"optional" json:"appInstallationId" yaml:"appInstallationId"`
	// Optional. SecretManager resource containing the private key of the GitHub App, formatted as 'projects/* /secrets/* /versions/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/developer_connect_connection#private_key_secret_version DeveloperConnectConnection#private_key_secret_version}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	PrivateKeySecretVersion *string `field:"optional" json:"privateKeySecretVersion" yaml:"privateKeySecretVersion"`
	// service_directory_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/developer_connect_connection#service_directory_config DeveloperConnectConnection#service_directory_config}
	ServiceDirectoryConfig *DeveloperConnectConnectionGithubEnterpriseConfigServiceDirectoryConfig `field:"optional" json:"serviceDirectoryConfig" yaml:"serviceDirectoryConfig"`
	// Optional. SSL certificate to use for requests to GitHub Enterprise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/developer_connect_connection#ssl_ca_certificate DeveloperConnectConnection#ssl_ca_certificate}
	SslCaCertificate *string `field:"optional" json:"sslCaCertificate" yaml:"sslCaCertificate"`
	// Optional. SecretManager resource containing the webhook secret of the GitHub App, formatted as 'projects/* /secrets/* /versions/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/developer_connect_connection#webhook_secret_secret_version DeveloperConnectConnection#webhook_secret_secret_version}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	WebhookSecretSecretVersion *string `field:"optional" json:"webhookSecretSecretVersion" yaml:"webhookSecretSecretVersion"`
}

