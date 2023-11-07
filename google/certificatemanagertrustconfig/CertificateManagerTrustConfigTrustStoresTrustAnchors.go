// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package certificatemanagertrustconfig


type CertificateManagerTrustConfigTrustStoresTrustAnchors struct {
	// PEM root certificate of the PKI used for validation. Each certificate provided in PEM format may occupy up to 5kB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.5.0/docs/resources/certificate_manager_trust_config#pem_certificate CertificateManagerTrustConfig#pem_certificate}
	PemCertificate *string `field:"optional" json:"pemCertificate" yaml:"pemCertificate"`
}

