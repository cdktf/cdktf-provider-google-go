// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection


type IntegrationConnectorsConnectionAuthConfigOauth2JwtBearerJwtClaims struct {
	// Value for the "aud" claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/integration_connectors_connection#audience IntegrationConnectorsConnection#audience}
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// Value for the "iss" claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/integration_connectors_connection#issuer IntegrationConnectorsConnection#issuer}
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// Value for the "sub" claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/integration_connectors_connection#subject IntegrationConnectorsConnection#subject}
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
}

