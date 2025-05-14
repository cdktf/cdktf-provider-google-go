// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package colabruntimetemplate


type ColabRuntimeTemplateMachineSpec struct {
	// The number of accelerators used by the runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/colab_runtime_template#accelerator_count ColabRuntimeTemplate#accelerator_count}
	AcceleratorCount *float64 `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// The type of hardware accelerator used by the runtime. If specified, acceleratorCount must also be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/colab_runtime_template#accelerator_type ColabRuntimeTemplate#accelerator_type}
	AcceleratorType *string `field:"optional" json:"acceleratorType" yaml:"acceleratorType"`
	// The Compute Engine machine type selected for the runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/colab_runtime_template#machine_type ColabRuntimeTemplate#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
}

