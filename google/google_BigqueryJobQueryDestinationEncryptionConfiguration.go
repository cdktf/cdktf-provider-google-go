// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type BigqueryJobQueryDestinationEncryptionConfiguration struct {
	// Describes the Cloud KMS encryption key that will be used to protect destination BigQuery table.
	//
	// The BigQuery Service Account associated with your project requires access to this encryption key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_job#kms_key_name BigqueryJob#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

