// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package certificatemanagercertificatemap


type CertificateManagerCertificateMapTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/certificate_manager_certificate_map#create CertificateManagerCertificateMap#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/certificate_manager_certificate_map#delete CertificateManagerCertificateMap#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/certificate_manager_certificate_map#update CertificateManagerCertificateMap#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

