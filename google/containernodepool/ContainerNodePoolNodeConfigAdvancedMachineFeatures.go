// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containernodepool


type ContainerNodePoolNodeConfigAdvancedMachineFeatures struct {
	// The number of threads per physical core.
	//
	// To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/container_node_pool#threads_per_core ContainerNodePool#threads_per_core}
	ThreadsPerCore *float64 `field:"required" json:"threadsPerCore" yaml:"threadsPerCore"`
	// Whether the node should have nested virtualization enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/container_node_pool#enable_nested_virtualization ContainerNodePool#enable_nested_virtualization}
	EnableNestedVirtualization interface{} `field:"optional" json:"enableNestedVirtualization" yaml:"enableNestedVirtualization"`
}

