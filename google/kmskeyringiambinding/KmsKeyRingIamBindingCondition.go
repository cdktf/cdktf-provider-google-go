package kmskeyringiambinding


type KmsKeyRingIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/kms_key_ring_iam_binding#expression KmsKeyRingIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/kms_key_ring_iam_binding#title KmsKeyRingIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/kms_key_ring_iam_binding#description KmsKeyRingIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

