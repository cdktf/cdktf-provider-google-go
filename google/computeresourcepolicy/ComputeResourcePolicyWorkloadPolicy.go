// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeresourcepolicy


type ComputeResourcePolicyWorkloadPolicy struct {
	// The type of workload policy. Possible values: ["HIGH_AVAILABILITY", "HIGH_THROUGHPUT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_resource_policy#type ComputeResourcePolicy#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The accelerator topology.
	//
	// This field can be set only when the workload policy type is HIGH_THROUGHPUT
	// and cannot be set if max topology distance is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_resource_policy#accelerator_topology ComputeResourcePolicy#accelerator_topology}
	AcceleratorTopology *string `field:"optional" json:"acceleratorTopology" yaml:"acceleratorTopology"`
	// The maximum topology distance.
	//
	// This field can be set only when the workload policy type is HIGH_THROUGHPUT
	// and cannot be set if accelerator topology is set. Possible values: ["BLOCK", "CLUSTER", "SUBBLOCK"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_resource_policy#max_topology_distance ComputeResourcePolicy#max_topology_distance}
	MaxTopologyDistance *string `field:"optional" json:"maxTopologyDistance" yaml:"maxTopologyDistance"`
}

