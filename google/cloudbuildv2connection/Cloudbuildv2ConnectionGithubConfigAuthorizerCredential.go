package cloudbuildv2connection


type Cloudbuildv2ConnectionGithubConfigAuthorizerCredential struct {
	// A SecretManager resource containing the OAuth token that authorizes the Cloud Build connection. Format: `projects/*\/secrets/*\/versions/*`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuildv2_connection#oauth_token_secret_version Cloudbuildv2Connection#oauth_token_secret_version}
	OauthTokenSecretVersion *string `field:"optional" json:"oauthTokenSecretVersion" yaml:"oauthTokenSecretVersion"`
}

