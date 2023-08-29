// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudiotregistryiambinding


type CloudiotRegistryIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.80.0/docs/resources/cloudiot_registry_iam_binding#expression CloudiotRegistryIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.80.0/docs/resources/cloudiot_registry_iam_binding#title CloudiotRegistryIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.80.0/docs/resources/cloudiot_registry_iam_binding#description CloudiotRegistryIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

