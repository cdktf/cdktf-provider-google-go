// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sccv2organizationsourceiambinding


type SccV2OrganizationSourceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/scc_v2_organization_source_iam_binding#expression SccV2OrganizationSourceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/scc_v2_organization_source_iam_binding#title SccV2OrganizationSourceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/scc_v2_organization_source_iam_binding#description SccV2OrganizationSourceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

