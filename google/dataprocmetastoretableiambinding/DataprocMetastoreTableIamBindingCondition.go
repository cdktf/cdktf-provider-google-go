// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocmetastoretableiambinding


type DataprocMetastoreTableIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/dataproc_metastore_table_iam_binding#expression DataprocMetastoreTableIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/dataproc_metastore_table_iam_binding#title DataprocMetastoreTableIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/dataproc_metastore_table_iam_binding#description DataprocMetastoreTableIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

