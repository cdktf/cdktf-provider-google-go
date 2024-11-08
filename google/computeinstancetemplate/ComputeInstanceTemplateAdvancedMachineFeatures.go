// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancetemplate


type ComputeInstanceTemplateAdvancedMachineFeatures struct {
	// Whether to enable nested virtualization or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/compute_instance_template#enable_nested_virtualization ComputeInstanceTemplate#enable_nested_virtualization}
	EnableNestedVirtualization interface{} `field:"optional" json:"enableNestedVirtualization" yaml:"enableNestedVirtualization"`
	// The number of threads per physical core.
	//
	// To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/compute_instance_template#threads_per_core ComputeInstanceTemplate#threads_per_core}
	ThreadsPerCore *float64 `field:"optional" json:"threadsPerCore" yaml:"threadsPerCore"`
	// Turbo frequency mode to use for the instance. Currently supported modes is "ALL_CORE_MAX".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/compute_instance_template#turbo_mode ComputeInstanceTemplate#turbo_mode}
	TurboMode *string `field:"optional" json:"turboMode" yaml:"turboMode"`
	// The number of physical cores to expose to an instance.
	//
	// Multiply by the number of threads per core to compute the total number of virtual CPUs to expose to the instance. If unset, the number of cores is inferred from the instance\'s nominal CPU count and the underlying platform\'s SMT width.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/compute_instance_template#visible_core_count ComputeInstanceTemplate#visible_core_count}
	VisibleCoreCount *float64 `field:"optional" json:"visibleCoreCount" yaml:"visibleCoreCount"`
}

