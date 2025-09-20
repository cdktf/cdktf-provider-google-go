// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containernodepool


type ContainerNodePoolNodeConfigConfidentialNodes struct {
	// Whether Confidential Nodes feature is enabled for all nodes in this pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/container_node_pool#enabled ContainerNodePool#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Defines the type of technology used by the confidential node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/container_node_pool#confidential_instance_type ContainerNodePool#confidential_instance_type}
	ConfidentialInstanceType *string `field:"optional" json:"confidentialInstanceType" yaml:"confidentialInstanceType"`
}

