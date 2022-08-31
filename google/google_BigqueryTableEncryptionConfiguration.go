// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type BigqueryTableEncryptionConfiguration struct {
	// The self link or full name of a key which should be used to encrypt this table.
	//
	// Note that the default bigquery service account will need to have encrypt/decrypt permissions on this key - you may want to see the google_bigquery_default_service_account datasource and the google_kms_crypto_key_iam_binding resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_table#kms_key_name BigqueryTable#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

