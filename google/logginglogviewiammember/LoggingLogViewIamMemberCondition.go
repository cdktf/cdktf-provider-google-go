// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logginglogviewiammember


type LoggingLogViewIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/logging_log_view_iam_member#expression LoggingLogViewIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/logging_log_view_iam_member#title LoggingLogViewIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/logging_log_view_iam_member#description LoggingLogViewIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

