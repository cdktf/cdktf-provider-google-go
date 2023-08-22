package healthcareconsentstoreiammember


type HealthcareConsentStoreIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/healthcare_consent_store_iam_member#expression HealthcareConsentStoreIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/healthcare_consent_store_iam_member#title HealthcareConsentStoreIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/healthcare_consent_store_iam_member#description HealthcareConsentStoreIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

