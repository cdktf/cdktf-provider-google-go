package alloydbbackup


type AlloydbBackupEncryptionConfig struct {
	// The fully-qualified resource name of the KMS key.
	//
	// Each Cloud KMS key is regionalized and has the following format: projects/[PROJECT]/locations/[REGION]/keyRings/[RING]/cryptoKeys/[KEY_NAME].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/alloydb_backup#kms_key_name AlloydbBackup#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
}

