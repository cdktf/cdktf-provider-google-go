// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagoogleiampolicy


type DataGoogleIamPolicyBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/data-sources/iam_policy#expression DataGoogleIamPolicy#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/data-sources/iam_policy#title DataGoogleIamPolicy#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/data-sources/iam_policy#description DataGoogleIamPolicy#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

