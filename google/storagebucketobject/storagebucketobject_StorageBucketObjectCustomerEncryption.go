package storagebucketobject


type StorageBucketObjectCustomerEncryption struct {
	// Base64 encoded customer supplied encryption key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_bucket_object#encryption_key StorageBucketObject#encryption_key}
	EncryptionKey *string `field:"required" json:"encryptionKey" yaml:"encryptionKey"`
	// The encryption algorithm. Default: AES256.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_bucket_object#encryption_algorithm StorageBucketObject#encryption_algorithm}
	EncryptionAlgorithm *string `field:"optional" json:"encryptionAlgorithm" yaml:"encryptionAlgorithm"`
}

