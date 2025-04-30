// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatecacapooliambinding


type PrivatecaCaPoolIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/privateca_ca_pool_iam_binding#expression PrivatecaCaPoolIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/privateca_ca_pool_iam_binding#title PrivatecaCaPoolIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/privateca_ca_pool_iam_binding#description PrivatecaCaPoolIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

