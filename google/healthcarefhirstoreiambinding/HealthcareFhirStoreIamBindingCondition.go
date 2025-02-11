// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcarefhirstoreiambinding


type HealthcareFhirStoreIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/healthcare_fhir_store_iam_binding#expression HealthcareFhirStoreIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/healthcare_fhir_store_iam_binding#title HealthcareFhirStoreIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/healthcare_fhir_store_iam_binding#description HealthcareFhirStoreIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

