// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocmetastoredatabaseiambinding


type DataprocMetastoreDatabaseIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/dataproc_metastore_database_iam_binding#expression DataprocMetastoreDatabaseIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/dataproc_metastore_database_iam_binding#title DataprocMetastoreDatabaseIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/dataproc_metastore_database_iam_binding#description DataprocMetastoreDatabaseIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

