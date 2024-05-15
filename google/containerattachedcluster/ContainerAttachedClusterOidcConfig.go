// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerattachedcluster


type ContainerAttachedClusterOidcConfig struct {
	// A JSON Web Token (JWT) issuer URI. 'issuer' must start with 'https://'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/container_attached_cluster#issuer_url ContainerAttachedCluster#issuer_url}
	IssuerUrl *string `field:"required" json:"issuerUrl" yaml:"issuerUrl"`
	// OIDC verification keys in JWKS format (RFC 7517).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/container_attached_cluster#jwks ContainerAttachedCluster#jwks}
	Jwks *string `field:"optional" json:"jwks" yaml:"jwks"`
}

