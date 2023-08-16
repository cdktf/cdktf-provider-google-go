package composerenvironment


type ComposerEnvironmentConfigEncryptionConfig struct {
	// Optional. Customer-managed Encryption Key available through Google's Key Management Service. Cannot be updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/composer_environment#kms_key_name ComposerEnvironment#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

