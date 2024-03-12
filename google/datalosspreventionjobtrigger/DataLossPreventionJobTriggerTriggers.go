// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerTriggers struct {
	// manual block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.20.0/docs/resources/data_loss_prevention_job_trigger#manual DataLossPreventionJobTrigger#manual}
	Manual *DataLossPreventionJobTriggerTriggersManual `field:"optional" json:"manual" yaml:"manual"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.20.0/docs/resources/data_loss_prevention_job_trigger#schedule DataLossPreventionJobTrigger#schedule}
	Schedule *DataLossPreventionJobTriggerTriggersSchedule `field:"optional" json:"schedule" yaml:"schedule"`
}

