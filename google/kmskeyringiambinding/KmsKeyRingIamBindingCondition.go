package kmskeyringiambinding


type KmsKeyRingIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/kms_key_ring_iam_binding#expression KmsKeyRingIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/kms_key_ring_iam_binding#title KmsKeyRingIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/kms_key_ring_iam_binding#description KmsKeyRingIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

