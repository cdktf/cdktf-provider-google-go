package storagetransferjob


type StorageTransferJobTransferSpecGcsDataSink struct {
	// Google Cloud Storage bucket name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_transfer_job#bucket_name StorageTransferJob#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Google Cloud Storage path in bucket to transfer.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_transfer_job#path StorageTransferJob#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

