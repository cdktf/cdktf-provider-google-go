// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type HealthcareFhirStoreIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/healthcare_fhir_store_iam_binding#expression HealthcareFhirStoreIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/healthcare_fhir_store_iam_binding#title HealthcareFhirStoreIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/healthcare_fhir_store_iam_binding#description HealthcareFhirStoreIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

