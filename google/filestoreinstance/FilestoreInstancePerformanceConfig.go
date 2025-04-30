// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package filestoreinstance


type FilestoreInstancePerformanceConfig struct {
	// fixed_iops block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/filestore_instance#fixed_iops FilestoreInstance#fixed_iops}
	FixedIops *FilestoreInstancePerformanceConfigFixedIops `field:"optional" json:"fixedIops" yaml:"fixedIops"`
	// iops_per_tb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/filestore_instance#iops_per_tb FilestoreInstance#iops_per_tb}
	IopsPerTb *FilestoreInstancePerformanceConfigIopsPerTb `field:"optional" json:"iopsPerTb" yaml:"iopsPerTb"`
}

