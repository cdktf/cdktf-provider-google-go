// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatecacertificateauthority


type PrivatecaCertificateAuthoritySubordinateConfigPemIssuerChain struct {
	// Expected to be in leaf-to-root order according to RFC 5246.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/privateca_certificate_authority#pem_certificates PrivatecaCertificateAuthority#pem_certificates}
	PemCertificates *[]*string `field:"optional" json:"pemCertificates" yaml:"pemCertificates"`
}

