// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatecacertificatetemplateiambinding


type PrivatecaCertificateTemplateIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/privateca_certificate_template_iam_binding#expression PrivatecaCertificateTemplateIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/privateca_certificate_template_iam_binding#title PrivatecaCertificateTemplateIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/privateca_certificate_template_iam_binding#description PrivatecaCertificateTemplateIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

