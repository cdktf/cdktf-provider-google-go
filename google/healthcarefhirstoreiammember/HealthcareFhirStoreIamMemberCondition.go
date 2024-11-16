// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcarefhirstoreiammember


type HealthcareFhirStoreIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/healthcare_fhir_store_iam_member#expression HealthcareFhirStoreIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/healthcare_fhir_store_iam_member#title HealthcareFhirStoreIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/healthcare_fhir_store_iam_member#description HealthcareFhirStoreIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

