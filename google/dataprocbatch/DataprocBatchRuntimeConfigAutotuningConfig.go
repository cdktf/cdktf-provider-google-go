// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocbatch


type DataprocBatchRuntimeConfigAutotuningConfig struct {
	// Optional. Scenarios for which tunings are applied. Possible values: ["SCALING", "BROADCAST_HASH_JOIN", "MEMORY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/dataproc_batch#scenarios DataprocBatch#scenarios}
	Scenarios *[]*string `field:"optional" json:"scenarios" yaml:"scenarios"`
}

