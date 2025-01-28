// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package colabruntimetemplateiammember


type ColabRuntimeTemplateIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/colab_runtime_template_iam_member#expression ColabRuntimeTemplateIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/colab_runtime_template_iam_member#title ColabRuntimeTemplateIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/colab_runtime_template_iam_member#description ColabRuntimeTemplateIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

