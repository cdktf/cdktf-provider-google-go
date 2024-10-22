// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamworkforcepoolprovider


type IamWorkforcePoolProviderExtraAttributesOauth2Client struct {
	// Represents the IdP and type of claims that should be fetched.
	//
	// * AZURE_AD_GROUPS_MAIL: Used to get the user's group claims from the Azure AD identity provider using configuration provided
	// in ExtraAttributesOAuth2Client and 'mail' property of the 'microsoft.graph.group' object is used for claim mapping.
	// See https://learn.microsoft.com/en-us/graph/api/resources/group?view=graph-rest-1.0#properties for more details on
	// 'microsoft.graph.group' properties. The attributes obtained from idntity provider are mapped to 'assertion.groups'. Possible values: ["AZURE_AD_GROUPS_MAIL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/iam_workforce_pool_provider#attributes_type IamWorkforcePoolProvider#attributes_type}
	AttributesType *string `field:"required" json:"attributesType" yaml:"attributesType"`
	// The OAuth 2.0 client ID for retrieving extra attributes from the identity provider. Required to get the Access Token using client credentials grant flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/iam_workforce_pool_provider#client_id IamWorkforcePoolProvider#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// client_secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/iam_workforce_pool_provider#client_secret IamWorkforcePoolProvider#client_secret}
	ClientSecret *IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecret `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// The OIDC identity provider's issuer URI.
	//
	// Must be a valid URI using the 'https' scheme. Required to get the OIDC discovery document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/iam_workforce_pool_provider#issuer_uri IamWorkforcePoolProvider#issuer_uri}
	IssuerUri *string `field:"required" json:"issuerUri" yaml:"issuerUri"`
	// query_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/iam_workforce_pool_provider#query_parameters IamWorkforcePoolProvider#query_parameters}
	QueryParameters *IamWorkforcePoolProviderExtraAttributesOauth2ClientQueryParameters `field:"optional" json:"queryParameters" yaml:"queryParameters"`
}

