package storagetransferjob


type StorageTransferJobTransferSpecHttpDataSource struct {
	// The URL that points to the file that stores the object list entries.
	//
	// This file must allow public access. Currently, only URLs with HTTP and HTTPS schemes are supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/storage_transfer_job#list_url StorageTransferJob#list_url}
	ListUrl *string `field:"required" json:"listUrl" yaml:"listUrl"`
}

