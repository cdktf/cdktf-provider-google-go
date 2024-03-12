// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatecacertificatetemplateiammember


type PrivatecaCertificateTemplateIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.20.0/docs/resources/privateca_certificate_template_iam_member#expression PrivatecaCertificateTemplateIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.20.0/docs/resources/privateca_certificate_template_iam_member#title PrivatecaCertificateTemplateIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.20.0/docs/resources/privateca_certificate_template_iam_member#description PrivatecaCertificateTemplateIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

