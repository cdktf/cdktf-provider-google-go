// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerazurenodepool


type ContainerAzureNodePoolMaxPodsConstraint struct {
	// The maximum number of pods to schedule on a single node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/container_azure_node_pool#max_pods_per_node ContainerAzureNodePool#max_pods_per_node}
	MaxPodsPerNode *float64 `field:"required" json:"maxPodsPerNode" yaml:"maxPodsPerNode"`
}

