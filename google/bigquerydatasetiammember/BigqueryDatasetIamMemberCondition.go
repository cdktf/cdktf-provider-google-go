// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerydatasetiammember


type BigqueryDatasetIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/bigquery_dataset_iam_member#expression BigqueryDatasetIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/bigquery_dataset_iam_member#title BigqueryDatasetIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/bigquery_dataset_iam_member#description BigqueryDatasetIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

