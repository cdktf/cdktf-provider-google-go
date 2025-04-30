// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregioninstancetemplate


type ComputeRegionInstanceTemplateAdvancedMachineFeatures struct {
	// Whether to enable nested virtualization or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/compute_region_instance_template#enable_nested_virtualization ComputeRegionInstanceTemplate#enable_nested_virtualization}
	EnableNestedVirtualization interface{} `field:"optional" json:"enableNestedVirtualization" yaml:"enableNestedVirtualization"`
	// Whether to enable UEFI networking or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/compute_region_instance_template#enable_uefi_networking ComputeRegionInstanceTemplate#enable_uefi_networking}
	EnableUefiNetworking interface{} `field:"optional" json:"enableUefiNetworking" yaml:"enableUefiNetworking"`
	// The PMU is a hardware component within the CPU core that monitors how the processor runs code.
	//
	// Valid values for the level of PMU are "STANDARD", "ENHANCED", and "ARCHITECTURAL".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/compute_region_instance_template#performance_monitoring_unit ComputeRegionInstanceTemplate#performance_monitoring_unit}
	PerformanceMonitoringUnit *string `field:"optional" json:"performanceMonitoringUnit" yaml:"performanceMonitoringUnit"`
	// The number of threads per physical core.
	//
	// To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/compute_region_instance_template#threads_per_core ComputeRegionInstanceTemplate#threads_per_core}
	ThreadsPerCore *float64 `field:"optional" json:"threadsPerCore" yaml:"threadsPerCore"`
	// Turbo frequency mode to use for the instance. Currently supported modes is "ALL_CORE_MAX".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/compute_region_instance_template#turbo_mode ComputeRegionInstanceTemplate#turbo_mode}
	TurboMode *string `field:"optional" json:"turboMode" yaml:"turboMode"`
	// The number of physical cores to expose to an instance.
	//
	// Multiply by the number of threads per core to compute the total number of virtual CPUs to expose to the instance. If unset, the number of cores is inferred from the instance\'s nominal CPU count and the underlying platform\'s SMT width.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/compute_region_instance_template#visible_core_count ComputeRegionInstanceTemplate#visible_core_count}
	VisibleCoreCount *float64 `field:"optional" json:"visibleCoreCount" yaml:"visibleCoreCount"`
}

