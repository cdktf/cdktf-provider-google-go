package iamworkloadidentitypoolprovider


type IamWorkloadIdentityPoolProviderOidc struct {
	// The OIDC issuer URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iam_workload_identity_pool_provider#issuer_uri IamWorkloadIdentityPoolProvider#issuer_uri}
	IssuerUri *string `field:"required" json:"issuerUri" yaml:"issuerUri"`
	// Acceptable values for the 'aud' field (audience) in the OIDC token.
	//
	// Token exchange
	// requests are rejected if the token audience does not match one of the configured
	// values. Each audience may be at most 256 characters. A maximum of 10 audiences may
	// be configured.
	//
	// If this list is empty, the OIDC token audience must be equal to the full canonical
	// resource name of the WorkloadIdentityPoolProvider, with or without the HTTPS prefix.
	// For example:
	// ```
	// //iam.googleapis.com/projects/<project-number>/locations/<location>/workloadIdentityPools/<pool-id>/providers/<provider-id>
	// https://iam.googleapis.com/projects/<project-number>/locations/<location>/workloadIdentityPools/<pool-id>/providers/<provider-id>
	// ```
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iam_workload_identity_pool_provider#allowed_audiences IamWorkloadIdentityPoolProvider#allowed_audiences}
	AllowedAudiences *[]*string `field:"optional" json:"allowedAudiences" yaml:"allowedAudiences"`
	// OIDC JWKs in JSON String format.
	//
	// For details on definition of a
	// JWK, see https:tools.ietf.org/html/rfc7517. If not set, then we
	// use the 'jwks_uri' from the discovery document fetched from the
	// .well-known path for the 'issuer_uri'. Currently, RSA and EC asymmetric
	// keys are supported. The JWK must use following format and include only
	// the following fields:
	// ```
	// {
	// "keys": [
	// {
	//       "kty": "RSA/EC",
	//       "alg": "<algorithm>",
	//       "use": "sig",
	//       "kid": "<key-id>",
	//       "n": "",
	//       "e": "",
	//       "x": "",
	//       "y": "",
	//       "crv": ""
	// }
	// ]
	// }
	// ```
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iam_workload_identity_pool_provider#jwks_json IamWorkloadIdentityPoolProvider#jwks_json}
	JwksJson *string `field:"optional" json:"jwksJson" yaml:"jwksJson"`
}

