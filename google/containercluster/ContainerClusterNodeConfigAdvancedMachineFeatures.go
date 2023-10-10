// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodeConfigAdvancedMachineFeatures struct {
	// The number of threads per physical core.
	//
	// To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.1.0/docs/resources/container_cluster#threads_per_core ContainerCluster#threads_per_core}
	ThreadsPerCore *float64 `field:"required" json:"threadsPerCore" yaml:"threadsPerCore"`
}

