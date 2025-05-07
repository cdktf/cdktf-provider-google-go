// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionsslcertificate


type ComputeRegionSslCertificateTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/compute_region_ssl_certificate#create ComputeRegionSslCertificate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/compute_region_ssl_certificate#delete ComputeRegionSslCertificate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

