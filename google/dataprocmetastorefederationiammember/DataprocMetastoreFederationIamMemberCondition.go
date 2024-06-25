// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocmetastorefederationiammember


type DataprocMetastoreFederationIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.35.0/docs/resources/dataproc_metastore_federation_iam_member#expression DataprocMetastoreFederationIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.35.0/docs/resources/dataproc_metastore_federation_iam_member#title DataprocMetastoreFederationIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.35.0/docs/resources/dataproc_metastore_federation_iam_member#description DataprocMetastoreFederationIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

