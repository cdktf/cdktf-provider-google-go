package cloudbuildv2connection


type Cloudbuildv2ConnectionGitlabConfig struct {
	// authorizer_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuildv2_connection#authorizer_credential Cloudbuildv2Connection#authorizer_credential}
	AuthorizerCredential *Cloudbuildv2ConnectionGitlabConfigAuthorizerCredential `field:"required" json:"authorizerCredential" yaml:"authorizerCredential"`
	// read_authorizer_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuildv2_connection#read_authorizer_credential Cloudbuildv2Connection#read_authorizer_credential}
	ReadAuthorizerCredential *Cloudbuildv2ConnectionGitlabConfigReadAuthorizerCredential `field:"required" json:"readAuthorizerCredential" yaml:"readAuthorizerCredential"`
	// Required. Immutable. SecretManager resource containing the webhook secret of a GitLab Enterprise project, formatted as `projects/*\/secrets/*\/versions/*`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuildv2_connection#webhook_secret_secret_version Cloudbuildv2Connection#webhook_secret_secret_version}
	WebhookSecretSecretVersion *string `field:"required" json:"webhookSecretSecretVersion" yaml:"webhookSecretSecretVersion"`
	// The URI of the GitLab Enterprise host this connection is for. If not specified, the default value is https://gitlab.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuildv2_connection#host_uri Cloudbuildv2Connection#host_uri}
	HostUri *string `field:"optional" json:"hostUri" yaml:"hostUri"`
	// service_directory_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuildv2_connection#service_directory_config Cloudbuildv2Connection#service_directory_config}
	ServiceDirectoryConfig *Cloudbuildv2ConnectionGitlabConfigServiceDirectoryConfig `field:"optional" json:"serviceDirectoryConfig" yaml:"serviceDirectoryConfig"`
	// SSL certificate to use for requests to GitLab Enterprise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuildv2_connection#ssl_ca Cloudbuildv2Connection#ssl_ca}
	SslCa *string `field:"optional" json:"sslCa" yaml:"sslCa"`
}

