package spannerdatabase


type SpannerDatabaseEncryptionConfig struct {
	// Fully qualified name of the KMS key to use to encrypt this database.
	//
	// This key must exist
	// in the same location as the Spanner Database.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/spanner_database#kms_key_name SpannerDatabase#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

