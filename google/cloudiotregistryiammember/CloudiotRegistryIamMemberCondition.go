// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudiotregistryiammember


type CloudiotRegistryIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.82.0/docs/resources/cloudiot_registry_iam_member#expression CloudiotRegistryIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.82.0/docs/resources/cloudiot_registry_iam_member#title CloudiotRegistryIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.82.0/docs/resources/cloudiot_registry_iam_member#description CloudiotRegistryIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

