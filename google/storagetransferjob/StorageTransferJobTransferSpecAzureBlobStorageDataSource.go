package storagetransferjob


type StorageTransferJobTransferSpecAzureBlobStorageDataSource struct {
	// azure_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/storage_transfer_job#azure_credentials StorageTransferJob#azure_credentials}
	AzureCredentials *StorageTransferJobTransferSpecAzureBlobStorageDataSourceAzureCredentials `field:"required" json:"azureCredentials" yaml:"azureCredentials"`
	// The container to transfer from the Azure Storage account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/storage_transfer_job#container StorageTransferJob#container}
	Container *string `field:"required" json:"container" yaml:"container"`
	// The name of the Azure Storage account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/storage_transfer_job#storage_account StorageTransferJob#storage_account}
	StorageAccount *string `field:"required" json:"storageAccount" yaml:"storageAccount"`
	// Root path to transfer objects.
	//
	// Must be an empty string or full path name that ends with a '/'. This field is treated as an object prefix. As such, it should generally not begin with a '/'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/storage_transfer_job#path StorageTransferJob#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

