// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containernodepool


type ContainerNodePoolPlacementPolicy struct {
	// Type defines the type of placement policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.0/docs/resources/container_node_pool#type ContainerNodePool#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// If set, refers to the name of a custom resource policy supplied by the user.
	//
	// The resource policy must be in the same project and region as the node pool. If not found, InvalidArgument error is returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.0/docs/resources/container_node_pool#policy_name ContainerNodePool#policy_name}
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// TPU placement topology for pod slice node pool. https://cloud.google.com/tpu/docs/types-topologies#tpu_topologies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.0/docs/resources/container_node_pool#tpu_topology ContainerNodePool#tpu_topology}
	TpuTopology *string `field:"optional" json:"tpuTopology" yaml:"tpuTopology"`
}

