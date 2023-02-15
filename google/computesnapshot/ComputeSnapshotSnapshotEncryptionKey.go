package computesnapshot


type ComputeSnapshotSnapshotEncryptionKey struct {
	// The name of the encryption key that is stored in Google Cloud KMS.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_snapshot#kms_key_self_link ComputeSnapshot#kms_key_self_link}
	KmsKeySelfLink *string `field:"optional" json:"kmsKeySelfLink" yaml:"kmsKeySelfLink"`
	// The service account used for the encryption request for the given KMS key.
	//
	// If absent, the Compute Engine Service Agent service account is used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_snapshot#kms_key_service_account ComputeSnapshot#kms_key_service_account}
	KmsKeyServiceAccount *string `field:"optional" json:"kmsKeyServiceAccount" yaml:"kmsKeyServiceAccount"`
	// Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_snapshot#raw_key ComputeSnapshot#raw_key}
	RawKey *string `field:"optional" json:"rawKey" yaml:"rawKey"`
}

