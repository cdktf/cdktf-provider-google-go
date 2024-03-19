// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcareconsentstoreiambinding


type HealthcareConsentStoreIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/healthcare_consent_store_iam_binding#expression HealthcareConsentStoreIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/healthcare_consent_store_iam_binding#title HealthcareConsentStoreIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/healthcare_consent_store_iam_binding#description HealthcareConsentStoreIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

