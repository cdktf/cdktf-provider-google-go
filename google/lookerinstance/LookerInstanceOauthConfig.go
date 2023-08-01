package lookerinstance


type LookerInstanceOauthConfig struct {
	// The client ID for the Oauth config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/looker_instance#client_id LookerInstance#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The client secret for the Oauth config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/looker_instance#client_secret LookerInstance#client_secret}
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
}

