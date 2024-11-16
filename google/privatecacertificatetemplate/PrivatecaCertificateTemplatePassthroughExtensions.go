// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatecacertificatetemplate


type PrivatecaCertificateTemplatePassthroughExtensions struct {
	// additional_extensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/privateca_certificate_template#additional_extensions PrivatecaCertificateTemplate#additional_extensions}
	AdditionalExtensions interface{} `field:"optional" json:"additionalExtensions" yaml:"additionalExtensions"`
	// Optional.
	//
	// A set of named X.509 extensions. Will be combined with additional_extensions to determine the full set of X.509 extensions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/privateca_certificate_template#known_extensions PrivatecaCertificateTemplate#known_extensions}
	KnownExtensions *[]*string `field:"optional" json:"knownExtensions" yaml:"knownExtensions"`
}

