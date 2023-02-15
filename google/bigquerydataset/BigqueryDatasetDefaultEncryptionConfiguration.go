package bigquerydataset


type BigqueryDatasetDefaultEncryptionConfiguration struct {
	// Describes the Cloud KMS encryption key that will be used to protect destination BigQuery table.
	//
	// The BigQuery Service Account associated with your project requires
	// access to this encryption key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_dataset#kms_key_name BigqueryDataset#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

