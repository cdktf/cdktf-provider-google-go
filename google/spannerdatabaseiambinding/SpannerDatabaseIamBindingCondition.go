// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spannerdatabaseiambinding


type SpannerDatabaseIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/spanner_database_iam_binding#expression SpannerDatabaseIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/spanner_database_iam_binding#title SpannerDatabaseIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/spanner_database_iam_binding#description SpannerDatabaseIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

