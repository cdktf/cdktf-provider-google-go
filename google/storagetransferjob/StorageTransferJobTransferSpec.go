package storagetransferjob


type StorageTransferJobTransferSpec struct {
	// aws_s3_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/storage_transfer_job#aws_s3_data_source StorageTransferJob#aws_s3_data_source}
	AwsS3DataSource *StorageTransferJobTransferSpecAwsS3DataSource `field:"optional" json:"awsS3DataSource" yaml:"awsS3DataSource"`
	// azure_blob_storage_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/storage_transfer_job#azure_blob_storage_data_source StorageTransferJob#azure_blob_storage_data_source}
	AzureBlobStorageDataSource *StorageTransferJobTransferSpecAzureBlobStorageDataSource `field:"optional" json:"azureBlobStorageDataSource" yaml:"azureBlobStorageDataSource"`
	// gcs_data_sink block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/storage_transfer_job#gcs_data_sink StorageTransferJob#gcs_data_sink}
	GcsDataSink *StorageTransferJobTransferSpecGcsDataSink `field:"optional" json:"gcsDataSink" yaml:"gcsDataSink"`
	// gcs_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/storage_transfer_job#gcs_data_source StorageTransferJob#gcs_data_source}
	GcsDataSource *StorageTransferJobTransferSpecGcsDataSource `field:"optional" json:"gcsDataSource" yaml:"gcsDataSource"`
	// http_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/storage_transfer_job#http_data_source StorageTransferJob#http_data_source}
	HttpDataSource *StorageTransferJobTransferSpecHttpDataSource `field:"optional" json:"httpDataSource" yaml:"httpDataSource"`
	// object_conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/storage_transfer_job#object_conditions StorageTransferJob#object_conditions}
	ObjectConditions *StorageTransferJobTransferSpecObjectConditions `field:"optional" json:"objectConditions" yaml:"objectConditions"`
	// posix_data_sink block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/storage_transfer_job#posix_data_sink StorageTransferJob#posix_data_sink}
	PosixDataSink *StorageTransferJobTransferSpecPosixDataSink `field:"optional" json:"posixDataSink" yaml:"posixDataSink"`
	// posix_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/storage_transfer_job#posix_data_source StorageTransferJob#posix_data_source}
	PosixDataSource *StorageTransferJobTransferSpecPosixDataSource `field:"optional" json:"posixDataSource" yaml:"posixDataSource"`
	// Specifies the agent pool name associated with the posix data source. When unspecified, the default name is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/storage_transfer_job#sink_agent_pool_name StorageTransferJob#sink_agent_pool_name}
	SinkAgentPoolName *string `field:"optional" json:"sinkAgentPoolName" yaml:"sinkAgentPoolName"`
	// Specifies the agent pool name associated with the posix data source. When unspecified, the default name is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/storage_transfer_job#source_agent_pool_name StorageTransferJob#source_agent_pool_name}
	SourceAgentPoolName *string `field:"optional" json:"sourceAgentPoolName" yaml:"sourceAgentPoolName"`
	// transfer_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/storage_transfer_job#transfer_options StorageTransferJob#transfer_options}
	TransferOptions *StorageTransferJobTransferSpecTransferOptions `field:"optional" json:"transferOptions" yaml:"transferOptions"`
}

