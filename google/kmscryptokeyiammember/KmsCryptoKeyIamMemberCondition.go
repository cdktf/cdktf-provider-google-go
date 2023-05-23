package kmscryptokeyiammember


type KmsCryptoKeyIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.66.0/docs/resources/kms_crypto_key_iam_member#expression KmsCryptoKeyIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.66.0/docs/resources/kms_crypto_key_iam_member#title KmsCryptoKeyIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.66.0/docs/resources/kms_crypto_key_iam_member#description KmsCryptoKeyIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

