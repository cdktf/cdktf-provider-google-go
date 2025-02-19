// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securesourcemanagerinstanceiambinding


type SecureSourceManagerInstanceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/secure_source_manager_instance_iam_binding#expression SecureSourceManagerInstanceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/secure_source_manager_instance_iam_binding#title SecureSourceManagerInstanceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/secure_source_manager_instance_iam_binding#description SecureSourceManagerInstanceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

