package secretmanagersecret


type SecretManagerSecretReplicationUserManagedReplicasCustomerManagedEncryption struct {
	// Describes the Cloud KMS encryption key that will be used to protect destination secret.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/secret_manager_secret#kms_key_name SecretManagerSecret#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

