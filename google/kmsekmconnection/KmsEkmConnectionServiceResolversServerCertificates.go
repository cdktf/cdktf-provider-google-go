// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kmsekmconnection


type KmsEkmConnectionServiceResolversServerCertificates struct {
	// Required. The raw certificate bytes in DER format. A base64-encoded string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/kms_ekm_connection#raw_der KmsEkmConnection#raw_der}
	RawDer *string `field:"required" json:"rawDer" yaml:"rawDer"`
	// Output only. The subject Alternative DNS names. Only present if parsed is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/kms_ekm_connection#subject_alternative_dns_names KmsEkmConnection#subject_alternative_dns_names}
	SubjectAlternativeDnsNames *[]*string `field:"optional" json:"subjectAlternativeDnsNames" yaml:"subjectAlternativeDnsNames"`
}

