// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath struct {
	// A url representing a file or path (no wildcards) in Cloud Storage. Example: 'gs://[BUCKET_NAME]/dictionary.txt'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/data_loss_prevention_job_trigger#path DataLossPreventionJobTrigger#path}
	Path *string `field:"required" json:"path" yaml:"path"`
}

