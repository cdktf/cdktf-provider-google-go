// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbuildv2connection


type Cloudbuildv2ConnectionBitbucketCloudConfig struct {
	// authorizer_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/cloudbuildv2_connection#authorizer_credential Cloudbuildv2Connection#authorizer_credential}
	AuthorizerCredential *Cloudbuildv2ConnectionBitbucketCloudConfigAuthorizerCredential `field:"required" json:"authorizerCredential" yaml:"authorizerCredential"`
	// read_authorizer_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/cloudbuildv2_connection#read_authorizer_credential Cloudbuildv2Connection#read_authorizer_credential}
	ReadAuthorizerCredential *Cloudbuildv2ConnectionBitbucketCloudConfigReadAuthorizerCredential `field:"required" json:"readAuthorizerCredential" yaml:"readAuthorizerCredential"`
	// Required. Immutable. SecretManager resource containing the webhook secret used to verify webhook events, formatted as 'projects/* /secrets/* /versions/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/cloudbuildv2_connection#webhook_secret_secret_version Cloudbuildv2Connection#webhook_secret_secret_version}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	WebhookSecretSecretVersion *string `field:"required" json:"webhookSecretSecretVersion" yaml:"webhookSecretSecretVersion"`
	// The Bitbucket Cloud Workspace ID to be connected to Google Cloud Platform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/cloudbuildv2_connection#workspace Cloudbuildv2Connection#workspace}
	Workspace *string `field:"required" json:"workspace" yaml:"workspace"`
}

