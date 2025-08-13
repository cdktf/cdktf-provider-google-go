// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagetransferjob


type StorageTransferJobTransferSpecAzureBlobStorageDataSource struct {
	// The container to transfer from the Azure Storage account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/storage_transfer_job#container StorageTransferJob#container}
	Container *string `field:"required" json:"container" yaml:"container"`
	// The name of the Azure Storage account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/storage_transfer_job#storage_account StorageTransferJob#storage_account}
	StorageAccount *string `field:"required" json:"storageAccount" yaml:"storageAccount"`
	// azure_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/storage_transfer_job#azure_credentials StorageTransferJob#azure_credentials}
	AzureCredentials *StorageTransferJobTransferSpecAzureBlobStorageDataSourceAzureCredentials `field:"optional" json:"azureCredentials" yaml:"azureCredentials"`
	// The Resource name of a secret in Secret Manager containing SAS Credentials in JSON form.
	//
	// Service Agent must have permissions to access secret. If credentials_secret is specified, do not specify azure_credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/storage_transfer_job#credentials_secret StorageTransferJob#credentials_secret}
	CredentialsSecret *string `field:"optional" json:"credentialsSecret" yaml:"credentialsSecret"`
	// federated_identity_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/storage_transfer_job#federated_identity_config StorageTransferJob#federated_identity_config}
	FederatedIdentityConfig *StorageTransferJobTransferSpecAzureBlobStorageDataSourceFederatedIdentityConfig `field:"optional" json:"federatedIdentityConfig" yaml:"federatedIdentityConfig"`
	// Root path to transfer objects.
	//
	// Must be an empty string or full path name that ends with a '/'. This field is treated as an object prefix. As such, it should generally not begin with a '/'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/storage_transfer_job#path StorageTransferJob#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

