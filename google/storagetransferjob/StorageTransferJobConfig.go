// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagetransferjob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageTransferJobConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Unique description to identify the Transfer Job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/storage_transfer_job#description StorageTransferJob#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// event_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/storage_transfer_job#event_stream StorageTransferJob#event_stream}
	EventStream *StorageTransferJobEventStream `field:"optional" json:"eventStream" yaml:"eventStream"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/storage_transfer_job#id StorageTransferJob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/storage_transfer_job#logging_config StorageTransferJob#logging_config}
	LoggingConfig *StorageTransferJobLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// The name of the Transfer Job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/storage_transfer_job#name StorageTransferJob#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// notification_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/storage_transfer_job#notification_config StorageTransferJob#notification_config}
	NotificationConfig *StorageTransferJobNotificationConfig `field:"optional" json:"notificationConfig" yaml:"notificationConfig"`
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/storage_transfer_job#project StorageTransferJob#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// replication_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/storage_transfer_job#replication_spec StorageTransferJob#replication_spec}
	ReplicationSpec *StorageTransferJobReplicationSpec `field:"optional" json:"replicationSpec" yaml:"replicationSpec"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/storage_transfer_job#schedule StorageTransferJob#schedule}
	Schedule *StorageTransferJobSchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Status of the job.
	//
	// Default: ENABLED. NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/storage_transfer_job#status StorageTransferJob#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// transfer_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/storage_transfer_job#transfer_spec StorageTransferJob#transfer_spec}
	TransferSpec *StorageTransferJobTransferSpec `field:"optional" json:"transferSpec" yaml:"transferSpec"`
}

