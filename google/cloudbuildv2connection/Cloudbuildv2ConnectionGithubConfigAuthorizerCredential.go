// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbuildv2connection


type Cloudbuildv2ConnectionGithubConfigAuthorizerCredential struct {
	// A SecretManager resource containing the OAuth token that authorizes the Cloud Build connection. Format: 'projects/* /secrets/* /versions/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/cloudbuildv2_connection#oauth_token_secret_version Cloudbuildv2Connection#oauth_token_secret_version}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	OauthTokenSecretVersion *string `field:"optional" json:"oauthTokenSecretVersion" yaml:"oauthTokenSecretVersion"`
}

