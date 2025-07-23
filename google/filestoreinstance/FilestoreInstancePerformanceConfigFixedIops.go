// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package filestoreinstance


type FilestoreInstancePerformanceConfigFixedIops struct {
	// The number of IOPS to provision for the instance. max_iops must be in multiple of 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/filestore_instance#max_iops FilestoreInstance#max_iops}
	MaxIops *float64 `field:"optional" json:"maxIops" yaml:"maxIops"`
}

