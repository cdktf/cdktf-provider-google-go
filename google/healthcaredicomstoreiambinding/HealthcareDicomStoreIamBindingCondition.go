// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcaredicomstoreiambinding


type HealthcareDicomStoreIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/healthcare_dicom_store_iam_binding#expression HealthcareDicomStoreIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/healthcare_dicom_store_iam_binding#title HealthcareDicomStoreIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/healthcare_dicom_store_iam_binding#description HealthcareDicomStoreIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

