// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package certificatemanagercertificateissuanceconfig


type CertificateManagerCertificateIssuanceConfigCertificateAuthorityConfigCertificateAuthorityServiceConfig struct {
	// A CA pool resource used to issue a certificate.
	//
	// The CA pool string has a relative resource path following the form
	// "projects/{project}/locations/{location}/caPools/{caPool}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/certificate_manager_certificate_issuance_config#ca_pool CertificateManagerCertificateIssuanceConfig#ca_pool}
	CaPool *string `field:"required" json:"caPool" yaml:"caPool"`
}

