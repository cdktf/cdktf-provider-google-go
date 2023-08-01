package storagetransferjob


type StorageTransferJobTransferSpecAwsS3DataSourceAwsAccessKey struct {
	// AWS Key ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/storage_transfer_job#access_key_id StorageTransferJob#access_key_id}
	AccessKeyId *string `field:"required" json:"accessKeyId" yaml:"accessKeyId"`
	// AWS Secret Access Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/storage_transfer_job#secret_access_key StorageTransferJob#secret_access_key}
	SecretAccessKey *string `field:"required" json:"secretAccessKey" yaml:"secretAccessKey"`
}

