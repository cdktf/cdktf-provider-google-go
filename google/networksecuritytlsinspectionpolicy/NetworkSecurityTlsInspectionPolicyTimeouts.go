// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecuritytlsinspectionpolicy


type NetworkSecurityTlsInspectionPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/network_security_tls_inspection_policy#create NetworkSecurityTlsInspectionPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/network_security_tls_inspection_policy#delete NetworkSecurityTlsInspectionPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/network_security_tls_inspection_policy#update NetworkSecurityTlsInspectionPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

