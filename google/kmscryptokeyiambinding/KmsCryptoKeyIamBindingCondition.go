package kmscryptokeyiambinding


type KmsCryptoKeyIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/kms_crypto_key_iam_binding#expression KmsCryptoKeyIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/kms_crypto_key_iam_binding#title KmsCryptoKeyIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/kms_crypto_key_iam_binding#description KmsCryptoKeyIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

