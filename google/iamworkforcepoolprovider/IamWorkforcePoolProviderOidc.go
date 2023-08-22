package iamworkforcepoolprovider


type IamWorkforcePoolProviderOidc struct {
	// The client ID. Must match the audience claim of the JWT issued by the identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/iam_workforce_pool_provider#client_id IamWorkforcePoolProvider#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The OIDC issuer URI. Must be a valid URI using the 'https' scheme.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/iam_workforce_pool_provider#issuer_uri IamWorkforcePoolProvider#issuer_uri}
	IssuerUri *string `field:"required" json:"issuerUri" yaml:"issuerUri"`
	// client_secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/iam_workforce_pool_provider#client_secret IamWorkforcePoolProvider#client_secret}
	ClientSecret *IamWorkforcePoolProviderOidcClientSecret `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// web_sso_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/iam_workforce_pool_provider#web_sso_config IamWorkforcePoolProvider#web_sso_config}
	WebSsoConfig *IamWorkforcePoolProviderOidcWebSsoConfig `field:"optional" json:"webSsoConfig" yaml:"webSsoConfig"`
}

