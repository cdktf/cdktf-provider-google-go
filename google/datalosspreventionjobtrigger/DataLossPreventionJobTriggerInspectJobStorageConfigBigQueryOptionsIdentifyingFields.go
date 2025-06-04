// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields struct {
	// Name of a BigQuery field to be returned with the findings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/data_loss_prevention_job_trigger#name DataLossPreventionJobTrigger#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

