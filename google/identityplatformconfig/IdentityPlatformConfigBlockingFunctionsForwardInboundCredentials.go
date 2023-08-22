package identityplatformconfig


type IdentityPlatformConfigBlockingFunctionsForwardInboundCredentials struct {
	// Whether to pass the user's OAuth identity provider's access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/identity_platform_config#access_token IdentityPlatformConfig#access_token}
	AccessToken interface{} `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Whether to pass the user's OIDC identity provider's ID token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/identity_platform_config#id_token IdentityPlatformConfig#id_token}
	IdToken interface{} `field:"optional" json:"idToken" yaml:"idToken"`
	// Whether to pass the user's OAuth identity provider's refresh token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/identity_platform_config#refresh_token IdentityPlatformConfig#refresh_token}
	RefreshToken interface{} `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

