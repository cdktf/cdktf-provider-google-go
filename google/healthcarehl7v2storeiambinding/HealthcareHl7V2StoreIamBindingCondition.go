package healthcarehl7v2storeiambinding


type HealthcareHl7V2StoreIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/healthcare_hl7_v2_store_iam_binding#expression HealthcareHl7V2StoreIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/healthcare_hl7_v2_store_iam_binding#title HealthcareHl7V2StoreIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/healthcare_hl7_v2_store_iam_binding#description HealthcareHl7V2StoreIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

