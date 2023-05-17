package alloydbcluster


type AlloydbClusterAutomatedBackupPolicyEncryptionConfig struct {
	// The fully-qualified resource name of the KMS key.
	//
	// Each Cloud KMS key is regionalized and has the following format: projects/[PROJECT]/locations/[REGION]/keyRings/[RING]/cryptoKeys/[KEY_NAME].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/alloydb_cluster#kms_key_name AlloydbCluster#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
}

