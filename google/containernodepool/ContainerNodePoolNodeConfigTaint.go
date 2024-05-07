// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containernodepool


type ContainerNodePoolNodeConfigTaint struct {
	// Effect for taint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.28.0/docs/resources/container_node_pool#effect ContainerNodePool#effect}
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Key for taint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.28.0/docs/resources/container_node_pool#key ContainerNodePool#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Value for taint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.28.0/docs/resources/container_node_pool#value ContainerNodePool#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

