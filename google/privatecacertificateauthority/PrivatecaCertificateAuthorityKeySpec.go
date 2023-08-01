package privatecacertificateauthority


type PrivatecaCertificateAuthorityKeySpec struct {
	// The algorithm to use for creating a managed Cloud KMS key for a for a simplified experience.
	//
	// All managed keys will be have their ProtectionLevel as HSM. Possible values: ["SIGN_HASH_ALGORITHM_UNSPECIFIED", "RSA_PSS_2048_SHA256", "RSA_PSS_3072_SHA256", "RSA_PSS_4096_SHA256", "RSA_PKCS1_2048_SHA256", "RSA_PKCS1_3072_SHA256", "RSA_PKCS1_4096_SHA256", "EC_P256_SHA256", "EC_P384_SHA384"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_certificate_authority#algorithm PrivatecaCertificateAuthority#algorithm}
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// The resource name for an existing Cloud KMS CryptoKeyVersion in the format 'projects/*\/locations/*\/keyRings/*\/cryptoKeys/*\/cryptoKeyVersions/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_certificate_authority#cloud_kms_key_version PrivatecaCertificateAuthority#cloud_kms_key_version}
	CloudKmsKeyVersion *string `field:"optional" json:"cloudKmsKeyVersion" yaml:"cloudKmsKeyVersion"`
}

