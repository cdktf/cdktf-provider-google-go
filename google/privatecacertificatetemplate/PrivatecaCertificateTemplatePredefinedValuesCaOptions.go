// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatecacertificatetemplate


type PrivatecaCertificateTemplatePredefinedValuesCaOptions struct {
	// Optional.
	//
	// Refers to the "CA" X.509 extension, which is a boolean value. When this value is true, the "CA" in Basic Constraints extension will be set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/privateca_certificate_template#is_ca PrivatecaCertificateTemplate#is_ca}
	IsCa interface{} `field:"optional" json:"isCa" yaml:"isCa"`
	// Optional.
	//
	// Refers to the "path length constraint" in Basic Constraints extension. For a CA certificate, this value describes the depth of
	// subordinate CA certificates that are allowed. If this value is less than 0, the request will fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/privateca_certificate_template#max_issuer_path_length PrivatecaCertificateTemplate#max_issuer_path_length}
	MaxIssuerPathLength *float64 `field:"optional" json:"maxIssuerPathLength" yaml:"maxIssuerPathLength"`
	// Optional.
	//
	// When true, the "CA" in Basic Constraints extension will be set to null and omitted from the CA certificate.
	// If both 'is_ca' and 'null_ca' are unset, the "CA" in Basic Constraints extension will be set to false.
	// Note that the behavior when 'is_ca = false' for this resource is different from the behavior in the Certificate Authority, Certificate and CaPool resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/privateca_certificate_template#null_ca PrivatecaCertificateTemplate#null_ca}
	NullCa interface{} `field:"optional" json:"nullCa" yaml:"nullCa"`
	// Optional.
	//
	// When true, the "path length constraint" in Basic Constraints extension will be set to 0.
	// if both 'max_issuer_path_length' and 'zero_max_issuer_path_length' are unset,
	// the max path length will be omitted from the CA certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/privateca_certificate_template#zero_max_issuer_path_length PrivatecaCertificateTemplate#zero_max_issuer_path_length}
	ZeroMaxIssuerPathLength interface{} `field:"optional" json:"zeroMaxIssuerPathLength" yaml:"zeroMaxIssuerPathLength"`
}

