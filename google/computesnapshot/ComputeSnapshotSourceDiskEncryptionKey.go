package computesnapshot


type ComputeSnapshotSourceDiskEncryptionKey struct {
	// The service account used for the encryption request for the given KMS key.
	//
	// If absent, the Compute Engine Service Agent service account is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.70.0/docs/resources/compute_snapshot#kms_key_service_account ComputeSnapshot#kms_key_service_account}
	KmsKeyServiceAccount *string `field:"optional" json:"kmsKeyServiceAccount" yaml:"kmsKeyServiceAccount"`
	// Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.70.0/docs/resources/compute_snapshot#raw_key ComputeSnapshot#raw_key}
	RawKey *string `field:"optional" json:"rawKey" yaml:"rawKey"`
}

