// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatecacertificateauthority


type PrivatecaCertificateAuthorityUserDefinedAccessUrls struct {
	// A list of URLs where this CertificateAuthority's CA certificate is published that is specified by users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/privateca_certificate_authority#aia_issuing_certificate_urls PrivatecaCertificateAuthority#aia_issuing_certificate_urls}
	AiaIssuingCertificateUrls *[]*string `field:"optional" json:"aiaIssuingCertificateUrls" yaml:"aiaIssuingCertificateUrls"`
	// A list of URLs where this CertificateAuthority's CRLs are published that is specified by users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/privateca_certificate_authority#crl_access_urls PrivatecaCertificateAuthority#crl_access_urls}
	CrlAccessUrls *[]*string `field:"optional" json:"crlAccessUrls" yaml:"crlAccessUrls"`
}

