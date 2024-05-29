// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamworkforcepoolprovider


type IamWorkforcePoolProviderOidc struct {
	// The client ID. Must match the audience claim of the JWT issued by the identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/iam_workforce_pool_provider#client_id IamWorkforcePoolProvider#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The OIDC issuer URI. Must be a valid URI using the 'https' scheme.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/iam_workforce_pool_provider#issuer_uri IamWorkforcePoolProvider#issuer_uri}
	IssuerUri *string `field:"required" json:"issuerUri" yaml:"issuerUri"`
	// client_secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/iam_workforce_pool_provider#client_secret IamWorkforcePoolProvider#client_secret}
	ClientSecret *IamWorkforcePoolProviderOidcClientSecret `field:"optional" json:"clientSecret" yaml:"clientSecret"`
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
	//   "keys": [
	//     {
	//           "kty": "RSA/EC",
	//           "alg": "<algorithm>",
	//           "use": "sig",
	//           "kid": "<key-id>",
	//           "n": "",
	//           "e": "",
	//           "x": "",
	//           "y": "",
	//           "crv": ""
	//     }
	//   ]
	// }
	// ```
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/iam_workforce_pool_provider#jwks_json IamWorkforcePoolProvider#jwks_json}
	JwksJson *string `field:"optional" json:"jwksJson" yaml:"jwksJson"`
	// web_sso_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/iam_workforce_pool_provider#web_sso_config IamWorkforcePoolProvider#web_sso_config}
	WebSsoConfig *IamWorkforcePoolProviderOidcWebSsoConfig `field:"optional" json:"webSsoConfig" yaml:"webSsoConfig"`
}

