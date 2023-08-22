package cloudbuildv2connection


type Cloudbuildv2ConnectionGitlabConfigReadAuthorizerCredential struct {
	// Required. A SecretManager resource containing the user token that authorizes the Cloud Build connection. Format: `projects/*\/secrets/*\/versions/*`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuildv2_connection#user_token_secret_version Cloudbuildv2Connection#user_token_secret_version}
	UserTokenSecretVersion *string `field:"required" json:"userTokenSecretVersion" yaml:"userTokenSecretVersion"`
}

