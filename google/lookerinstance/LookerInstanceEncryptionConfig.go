package lookerinstance


type LookerInstanceEncryptionConfig struct {
	// Name of the customer managed encryption key (CMEK) in KMS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/looker_instance#kms_key_name LookerInstance#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
}

