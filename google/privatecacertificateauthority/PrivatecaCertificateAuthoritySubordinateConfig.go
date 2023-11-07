// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatecacertificateauthority


type PrivatecaCertificateAuthoritySubordinateConfig struct {
	// This can refer to a CertificateAuthority that was used to create a subordinate CertificateAuthority.
	//
	// This field is used for information
	// and usability purposes only. The resource name is in the format
	// 'projects/* /locations/* /caPools/* /certificateAuthorities/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.5.0/docs/resources/privateca_certificate_authority#certificate_authority PrivatecaCertificateAuthority#certificate_authority}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	CertificateAuthority *string `field:"optional" json:"certificateAuthority" yaml:"certificateAuthority"`
	// pem_issuer_chain block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.5.0/docs/resources/privateca_certificate_authority#pem_issuer_chain PrivatecaCertificateAuthority#pem_issuer_chain}
	PemIssuerChain *PrivatecaCertificateAuthoritySubordinateConfigPemIssuerChain `field:"optional" json:"pemIssuerChain" yaml:"pemIssuerChain"`
}

