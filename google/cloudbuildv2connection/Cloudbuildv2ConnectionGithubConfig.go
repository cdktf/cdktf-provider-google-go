package cloudbuildv2connection


type Cloudbuildv2ConnectionGithubConfig struct {
	// GitHub App installation id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuildv2_connection#app_installation_id Cloudbuildv2Connection#app_installation_id}
	AppInstallationId *float64 `field:"optional" json:"appInstallationId" yaml:"appInstallationId"`
	// authorizer_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuildv2_connection#authorizer_credential Cloudbuildv2Connection#authorizer_credential}
	AuthorizerCredential *Cloudbuildv2ConnectionGithubConfigAuthorizerCredential `field:"optional" json:"authorizerCredential" yaml:"authorizerCredential"`
}

