// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeploycustomtargettypeiambinding


type ClouddeployCustomTargetTypeIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/clouddeploy_custom_target_type_iam_binding#expression ClouddeployCustomTargetTypeIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/clouddeploy_custom_target_type_iam_binding#title ClouddeployCustomTargetTypeIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/clouddeploy_custom_target_type_iam_binding#description ClouddeployCustomTargetTypeIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

