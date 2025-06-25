// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagetransferjob


type StorageTransferJobLoggingConfig struct {
	// For transfers with a PosixFilesystem source, this option enables the Cloud Storage transfer logs for this transfer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/storage_transfer_job#enable_on_prem_gcs_transfer_logs StorageTransferJob#enable_on_prem_gcs_transfer_logs}
	EnableOnPremGcsTransferLogs interface{} `field:"optional" json:"enableOnPremGcsTransferLogs" yaml:"enableOnPremGcsTransferLogs"`
	// Specifies the actions to be logged. Not supported for transfers with PosifxFilesystem data sources; use enable_on_prem_gcs_transfer_logs instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/storage_transfer_job#log_actions StorageTransferJob#log_actions}
	LogActions *[]*string `field:"optional" json:"logActions" yaml:"logActions"`
	// States in which logActions are logged. Not supported for transfers with PosifxFilesystem data sources; use enable_on_prem_gcs_transfer_logs instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/storage_transfer_job#log_action_states StorageTransferJob#log_action_states}
	LogActionStates *[]*string `field:"optional" json:"logActionStates" yaml:"logActionStates"`
}

