// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcaredatasetiambinding


type HealthcareDatasetIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/healthcare_dataset_iam_binding#expression HealthcareDatasetIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/healthcare_dataset_iam_binding#title HealthcareDatasetIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/healthcare_dataset_iam_binding#description HealthcareDatasetIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

