package healthcaredicomstoreiambinding


type HealthcareDicomStoreIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/healthcare_dicom_store_iam_binding#expression HealthcareDicomStoreIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/healthcare_dicom_store_iam_binding#title HealthcareDicomStoreIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/healthcare_dicom_store_iam_binding#description HealthcareDicomStoreIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

