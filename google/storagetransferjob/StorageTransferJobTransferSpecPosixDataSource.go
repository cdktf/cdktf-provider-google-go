package storagetransferjob


type StorageTransferJobTransferSpecPosixDataSource struct {
	// Root directory path to the filesystem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.63.1/docs/resources/storage_transfer_job#root_directory StorageTransferJob#root_directory}
	RootDirectory *string `field:"required" json:"rootDirectory" yaml:"rootDirectory"`
}

