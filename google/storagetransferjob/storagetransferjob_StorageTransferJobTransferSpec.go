package storagetransferjob


type StorageTransferJobTransferSpec struct {
	// aws_s3_data_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_transfer_job#aws_s3_data_source StorageTransferJob#aws_s3_data_source}
	AwsS3DataSource *StorageTransferJobTransferSpecAwsS3DataSource `field:"optional" json:"awsS3DataSource" yaml:"awsS3DataSource"`
	// azure_blob_storage_data_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_transfer_job#azure_blob_storage_data_source StorageTransferJob#azure_blob_storage_data_source}
	AzureBlobStorageDataSource *StorageTransferJobTransferSpecAzureBlobStorageDataSource `field:"optional" json:"azureBlobStorageDataSource" yaml:"azureBlobStorageDataSource"`
	// gcs_data_sink block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_transfer_job#gcs_data_sink StorageTransferJob#gcs_data_sink}
	GcsDataSink *StorageTransferJobTransferSpecGcsDataSink `field:"optional" json:"gcsDataSink" yaml:"gcsDataSink"`
	// gcs_data_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_transfer_job#gcs_data_source StorageTransferJob#gcs_data_source}
	GcsDataSource *StorageTransferJobTransferSpecGcsDataSource `field:"optional" json:"gcsDataSource" yaml:"gcsDataSource"`
	// http_data_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_transfer_job#http_data_source StorageTransferJob#http_data_source}
	HttpDataSource *StorageTransferJobTransferSpecHttpDataSource `field:"optional" json:"httpDataSource" yaml:"httpDataSource"`
	// object_conditions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_transfer_job#object_conditions StorageTransferJob#object_conditions}
	ObjectConditions *StorageTransferJobTransferSpecObjectConditions `field:"optional" json:"objectConditions" yaml:"objectConditions"`
	// posix_data_sink block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_transfer_job#posix_data_sink StorageTransferJob#posix_data_sink}
	PosixDataSink *StorageTransferJobTransferSpecPosixDataSink `field:"optional" json:"posixDataSink" yaml:"posixDataSink"`
	// posix_data_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_transfer_job#posix_data_source StorageTransferJob#posix_data_source}
	PosixDataSource *StorageTransferJobTransferSpecPosixDataSource `field:"optional" json:"posixDataSource" yaml:"posixDataSource"`
	// transfer_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_transfer_job#transfer_options StorageTransferJob#transfer_options}
	TransferOptions *StorageTransferJobTransferSpecTransferOptions `field:"optional" json:"transferOptions" yaml:"transferOptions"`
}

