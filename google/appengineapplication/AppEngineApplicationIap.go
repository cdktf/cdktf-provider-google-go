package appengineapplication


type AppEngineApplicationIap struct {
	// OAuth2 client ID to use for the authentication flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/app_engine_application#oauth2_client_id AppEngineApplication#oauth2_client_id}
	Oauth2ClientId *string `field:"required" json:"oauth2ClientId" yaml:"oauth2ClientId"`
	// OAuth2 client secret to use for the authentication flow.
	//
	// The SHA-256 hash of the value is returned in the oauth2ClientSecretSha256 field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/app_engine_application#oauth2_client_secret AppEngineApplication#oauth2_client_secret}
	Oauth2ClientSecret *string `field:"required" json:"oauth2ClientSecret" yaml:"oauth2ClientSecret"`
	// Adapted for use with the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/app_engine_application#enabled AppEngineApplication#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

