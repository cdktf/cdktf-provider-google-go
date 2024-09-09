// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeploytargetiambinding


type ClouddeployTargetIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/clouddeploy_target_iam_binding#expression ClouddeployTargetIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/clouddeploy_target_iam_binding#title ClouddeployTargetIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/clouddeploy_target_iam_binding#description ClouddeployTargetIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

