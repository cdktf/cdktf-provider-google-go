// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alloydbinstance


type AlloydbInstanceMachineConfig struct {
	// The number of CPU's in the VM instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/alloydb_instance#cpu_count AlloydbInstance#cpu_count}
	CpuCount *float64 `field:"optional" json:"cpuCount" yaml:"cpuCount"`
}

