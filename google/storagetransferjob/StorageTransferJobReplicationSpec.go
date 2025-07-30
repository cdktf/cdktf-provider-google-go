// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagetransferjob


type StorageTransferJobReplicationSpec struct {
	// gcs_data_sink block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/storage_transfer_job#gcs_data_sink StorageTransferJob#gcs_data_sink}
	GcsDataSink *StorageTransferJobReplicationSpecGcsDataSink `field:"optional" json:"gcsDataSink" yaml:"gcsDataSink"`
	// gcs_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/storage_transfer_job#gcs_data_source StorageTransferJob#gcs_data_source}
	GcsDataSource *StorageTransferJobReplicationSpecGcsDataSource `field:"optional" json:"gcsDataSource" yaml:"gcsDataSource"`
	// object_conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/storage_transfer_job#object_conditions StorageTransferJob#object_conditions}
	ObjectConditions *StorageTransferJobReplicationSpecObjectConditions `field:"optional" json:"objectConditions" yaml:"objectConditions"`
	// transfer_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/storage_transfer_job#transfer_options StorageTransferJob#transfer_options}
	TransferOptions *StorageTransferJobReplicationSpecTransferOptions `field:"optional" json:"transferOptions" yaml:"transferOptions"`
}

