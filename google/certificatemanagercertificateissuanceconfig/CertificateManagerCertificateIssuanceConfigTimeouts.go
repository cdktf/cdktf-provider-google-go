// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package certificatemanagercertificateissuanceconfig


type CertificateManagerCertificateIssuanceConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/certificate_manager_certificate_issuance_config#create CertificateManagerCertificateIssuanceConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/certificate_manager_certificate_issuance_config#delete CertificateManagerCertificateIssuanceConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/certificate_manager_certificate_issuance_config#update CertificateManagerCertificateIssuanceConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

