package spannerdatabase


type SpannerDatabaseEncryptionConfig struct {
	// Fully qualified name of the KMS key to use to encrypt this database.
	//
	// This key must exist
	// in the same location as the Spanner Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/spanner_database#kms_key_name SpannerDatabase#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

