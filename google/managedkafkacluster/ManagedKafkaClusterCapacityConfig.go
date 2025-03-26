// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedkafkacluster


type ManagedKafkaClusterCapacityConfig struct {
	// The memory to provision for the cluster in bytes.
	//
	// The value must be between 1 GiB and 8 GiB per vCPU. Ex. 1024Mi, 4Gi.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/managed_kafka_cluster#memory_bytes ManagedKafkaCluster#memory_bytes}
	MemoryBytes *string `field:"required" json:"memoryBytes" yaml:"memoryBytes"`
	// The number of vCPUs to provision for the cluster. The minimum is 3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/managed_kafka_cluster#vcpu_count ManagedKafkaCluster#vcpu_count}
	VcpuCount *string `field:"required" json:"vcpuCount" yaml:"vcpuCount"`
}

