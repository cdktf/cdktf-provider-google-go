// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptions struct {
	// kind block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/data_loss_prevention_job_trigger#kind DataLossPreventionJobTrigger#kind}
	Kind *DataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptionsKind `field:"required" json:"kind" yaml:"kind"`
	// partition_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/data_loss_prevention_job_trigger#partition_id DataLossPreventionJobTrigger#partition_id}
	PartitionId *DataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId `field:"required" json:"partitionId" yaml:"partitionId"`
}

