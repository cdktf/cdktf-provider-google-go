package iamworkforcepoolprovider


type IamWorkforcePoolProviderOidcWebSsoConfig struct {
	// The behavior for how OIDC Claims are included in the 'assertion' object used for attribute mapping and attribute condition.
	//
	// ONLY_ID_TOKEN_CLAIMS: Only include ID Token Claims. Possible values: ["ONLY_ID_TOKEN_CLAIMS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/iam_workforce_pool_provider#assertion_claims_behavior IamWorkforcePoolProvider#assertion_claims_behavior}
	AssertionClaimsBehavior *string `field:"required" json:"assertionClaimsBehavior" yaml:"assertionClaimsBehavior"`
	// The Response Type to request for in the OIDC Authorization Request for web sign-in.
	//
	// ID_TOKEN: The 'response_type=id_token' selection uses the Implicit Flow for web sign-in. Possible values: ["ID_TOKEN"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/iam_workforce_pool_provider#response_type IamWorkforcePoolProvider#response_type}
	ResponseType *string `field:"required" json:"responseType" yaml:"responseType"`
}

