// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfunctionsfunctioniammember


type CloudfunctionsFunctionIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/cloudfunctions_function_iam_member#expression CloudfunctionsFunctionIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/cloudfunctions_function_iam_member#title CloudfunctionsFunctionIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/cloudfunctions_function_iam_member#description CloudfunctionsFunctionIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

