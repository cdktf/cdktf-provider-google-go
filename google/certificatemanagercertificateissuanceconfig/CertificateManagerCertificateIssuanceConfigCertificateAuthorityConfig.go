// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package certificatemanagercertificateissuanceconfig


type CertificateManagerCertificateIssuanceConfigCertificateAuthorityConfig struct {
	// certificate_authority_service_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.3.0/docs/resources/certificate_manager_certificate_issuance_config#certificate_authority_service_config CertificateManagerCertificateIssuanceConfig#certificate_authority_service_config}
	CertificateAuthorityServiceConfig *CertificateManagerCertificateIssuanceConfigCertificateAuthorityConfigCertificateAuthorityServiceConfig `field:"optional" json:"certificateAuthorityServiceConfig" yaml:"certificateAuthorityServiceConfig"`
}

