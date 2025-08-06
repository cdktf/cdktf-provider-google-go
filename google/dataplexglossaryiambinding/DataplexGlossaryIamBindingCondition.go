// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexglossaryiambinding


type DataplexGlossaryIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dataplex_glossary_iam_binding#expression DataplexGlossaryIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dataplex_glossary_iam_binding#title DataplexGlossaryIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dataplex_glossary_iam_binding#description DataplexGlossaryIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

