// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplextaskiambinding


type DataplexTaskIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.83.0/docs/resources/dataplex_task_iam_binding#expression DataplexTaskIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.83.0/docs/resources/dataplex_task_iam_binding#title DataplexTaskIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.83.0/docs/resources/dataplex_task_iam_binding#description DataplexTaskIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

