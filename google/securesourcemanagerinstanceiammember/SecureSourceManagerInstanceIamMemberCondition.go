// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securesourcemanagerinstanceiammember


type SecureSourceManagerInstanceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/secure_source_manager_instance_iam_member#expression SecureSourceManagerInstanceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/secure_source_manager_instance_iam_member#title SecureSourceManagerInstanceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/secure_source_manager_instance_iam_member#description SecureSourceManagerInstanceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

