// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocmetastoredatabaseiammember


type DataprocMetastoreDatabaseIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/dataproc_metastore_database_iam_member#expression DataprocMetastoreDatabaseIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/dataproc_metastore_database_iam_member#title DataprocMetastoreDatabaseIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/dataproc_metastore_database_iam_member#description DataprocMetastoreDatabaseIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

