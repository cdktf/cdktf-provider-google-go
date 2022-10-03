package healthcarehl7v2storeiammember


type HealthcareHl7V2StoreIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/healthcare_hl7_v2_store_iam_member#expression HealthcareHl7V2StoreIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/healthcare_hl7_v2_store_iam_member#title HealthcareHl7V2StoreIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/healthcare_hl7_v2_store_iam_member#description HealthcareHl7V2StoreIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

