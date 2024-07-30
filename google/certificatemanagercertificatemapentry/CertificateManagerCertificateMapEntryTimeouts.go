// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package certificatemanagercertificatemapentry


type CertificateManagerCertificateMapEntryTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/certificate_manager_certificate_map_entry#create CertificateManagerCertificateMapEntry#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/certificate_manager_certificate_map_entry#delete CertificateManagerCertificateMapEntry#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/certificate_manager_certificate_map_entry#update CertificateManagerCertificateMapEntry#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

