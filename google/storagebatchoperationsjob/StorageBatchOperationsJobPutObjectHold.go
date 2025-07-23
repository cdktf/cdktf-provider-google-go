// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagebatchoperationsjob


type StorageBatchOperationsJobPutObjectHold struct {
	// set/unset to update event based hold for objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/storage_batch_operations_job#event_based_hold StorageBatchOperationsJob#event_based_hold}
	EventBasedHold *string `field:"optional" json:"eventBasedHold" yaml:"eventBasedHold"`
	// set/unset to update temporary based hold for objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/storage_batch_operations_job#temporary_hold StorageBatchOperationsJob#temporary_hold}
	TemporaryHold *string `field:"optional" json:"temporaryHold" yaml:"temporaryHold"`
}

