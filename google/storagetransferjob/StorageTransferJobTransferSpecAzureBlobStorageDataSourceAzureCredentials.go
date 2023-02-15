package storagetransferjob


type StorageTransferJobTransferSpecAzureBlobStorageDataSourceAzureCredentials struct {
	// Azure shared access signature.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_transfer_job#sas_token StorageTransferJob#sas_token}
	SasToken *string `field:"required" json:"sasToken" yaml:"sasToken"`
}

