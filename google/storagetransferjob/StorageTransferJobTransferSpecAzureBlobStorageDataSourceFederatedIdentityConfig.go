// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagetransferjob


type StorageTransferJobTransferSpecAzureBlobStorageDataSourceFederatedIdentityConfig struct {
	// The client (application) ID of the application with federated credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/storage_transfer_job#client_id StorageTransferJob#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The tenant (directory) ID of the application with federated credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/storage_transfer_job#tenant_id StorageTransferJob#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
}

