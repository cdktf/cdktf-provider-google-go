// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerydatasetiambinding


type BigqueryDatasetIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/bigquery_dataset_iam_binding#expression BigqueryDatasetIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/bigquery_dataset_iam_binding#title BigqueryDatasetIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/bigquery_dataset_iam_binding#description BigqueryDatasetIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

