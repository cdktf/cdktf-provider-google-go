// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatecacertificate


type PrivatecaCertificateConfigA struct {
	// public_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/privateca_certificate#public_key PrivatecaCertificate#public_key}
	PublicKey *PrivatecaCertificateConfigPublicKey `field:"required" json:"publicKey" yaml:"publicKey"`
	// subject_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/privateca_certificate#subject_config PrivatecaCertificate#subject_config}
	SubjectConfig *PrivatecaCertificateConfigSubjectConfig `field:"required" json:"subjectConfig" yaml:"subjectConfig"`
	// x509_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/privateca_certificate#x509_config PrivatecaCertificate#x509_config}
	X509Config *PrivatecaCertificateConfigX509Config `field:"required" json:"x509Config" yaml:"x509Config"`
	// subject_key_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/privateca_certificate#subject_key_id PrivatecaCertificate#subject_key_id}
	SubjectKeyId *PrivatecaCertificateConfigSubjectKeyId `field:"optional" json:"subjectKeyId" yaml:"subjectKeyId"`
}

