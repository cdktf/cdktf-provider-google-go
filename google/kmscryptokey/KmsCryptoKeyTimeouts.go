package kmscryptokey


type KmsCryptoKeyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/kms_crypto_key#create KmsCryptoKey#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/kms_crypto_key#delete KmsCryptoKey#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/kms_crypto_key#update KmsCryptoKey#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

