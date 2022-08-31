// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type StorageTransferJobTransferSpecPosixDataSource struct {
	// Root directory path to the filesystem.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_transfer_job#root_directory StorageTransferJob#root_directory}
	RootDirectory *string `field:"required" json:"rootDirectory" yaml:"rootDirectory"`
}

