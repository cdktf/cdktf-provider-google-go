package iamworkforcepoolprovider


type IamWorkforcePoolProviderOidc struct {
	// The client ID. Must match the audience claim of the JWT issued by the identity provider.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iam_workforce_pool_provider#client_id IamWorkforcePoolProvider#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The OIDC issuer URI. Must be a valid URI using the 'https' scheme.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iam_workforce_pool_provider#issuer_uri IamWorkforcePoolProvider#issuer_uri}
	IssuerUri *string `field:"required" json:"issuerUri" yaml:"issuerUri"`
}

