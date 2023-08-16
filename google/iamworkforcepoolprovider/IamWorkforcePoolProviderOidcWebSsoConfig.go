package iamworkforcepoolprovider


type IamWorkforcePoolProviderOidcWebSsoConfig struct {
	// The behavior for how OIDC Claims are included in the 'assertion' object used for attribute mapping and attribute condition.
	//
	// MERGE_USER_INFO_OVER_ID_TOKEN_CLAIMS: Merge the UserInfo Endpoint Claims with ID Token Claims, preferring UserInfo Claim Values for the same Claim Name. This option is available only for the Authorization Code Flow.
	// ONLY_ID_TOKEN_CLAIMS: Only include ID Token Claims. Possible values: ["MERGE_USER_INFO_OVER_ID_TOKEN_CLAIMS", "ONLY_ID_TOKEN_CLAIMS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iam_workforce_pool_provider#assertion_claims_behavior IamWorkforcePoolProvider#assertion_claims_behavior}
	AssertionClaimsBehavior *string `field:"required" json:"assertionClaimsBehavior" yaml:"assertionClaimsBehavior"`
	// The Response Type to request for in the OIDC Authorization Request for web sign-in.
	//
	// The 'CODE' Response Type is recommended to avoid the Implicit Flow, for security reasons.
	// CODE: The 'response_type=code' selection uses the Authorization Code Flow for web sign-in. Requires a configured client secret.
	// ID_TOKEN: The 'response_type=id_token' selection uses the Implicit Flow for web sign-in. Possible values: ["CODE", "ID_TOKEN"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iam_workforce_pool_provider#response_type IamWorkforcePoolProvider#response_type}
	ResponseType *string `field:"required" json:"responseType" yaml:"responseType"`
}

