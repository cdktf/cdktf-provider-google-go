// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocclusteriambinding


type DataprocClusterIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.11.0/docs/resources/dataproc_cluster_iam_binding#expression DataprocClusterIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.11.0/docs/resources/dataproc_cluster_iam_binding#title DataprocClusterIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.11.0/docs/resources/dataproc_cluster_iam_binding#description DataprocClusterIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

