// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computesubnetworkiambinding


type ComputeSubnetworkIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/compute_subnetwork_iam_binding#expression ComputeSubnetworkIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/compute_subnetwork_iam_binding#title ComputeSubnetworkIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/compute_subnetwork_iam_binding#description ComputeSubnetworkIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

