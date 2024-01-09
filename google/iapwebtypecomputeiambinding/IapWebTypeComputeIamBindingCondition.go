// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iapwebtypecomputeiambinding


type IapWebTypeComputeIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.11.0/docs/resources/iap_web_type_compute_iam_binding#expression IapWebTypeComputeIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.11.0/docs/resources/iap_web_type_compute_iam_binding#title IapWebTypeComputeIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.11.0/docs/resources/iap_web_type_compute_iam_binding#description IapWebTypeComputeIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

