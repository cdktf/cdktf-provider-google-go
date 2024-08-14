// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexdatascan


type DataplexDatascanExecutionSpecTrigger struct {
	// on_demand block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.41.0/docs/resources/dataplex_datascan#on_demand DataplexDatascan#on_demand}
	OnDemand *DataplexDatascanExecutionSpecTriggerOnDemand `field:"optional" json:"onDemand" yaml:"onDemand"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.41.0/docs/resources/dataplex_datascan#schedule DataplexDatascan#schedule}
	Schedule *DataplexDatascanExecutionSpecTriggerSchedule `field:"optional" json:"schedule" yaml:"schedule"`
}

