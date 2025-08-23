// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcaredicomstoreiammember


type HealthcareDicomStoreIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/healthcare_dicom_store_iam_member#expression HealthcareDicomStoreIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/healthcare_dicom_store_iam_member#title HealthcareDicomStoreIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/healthcare_dicom_store_iam_member#description HealthcareDicomStoreIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

