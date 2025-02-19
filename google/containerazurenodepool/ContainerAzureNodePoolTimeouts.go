// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerazurenodepool


type ContainerAzureNodePoolTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/container_azure_node_pool#create ContainerAzureNodePool#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/container_azure_node_pool#delete ContainerAzureNodePool#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/container_azure_node_pool#update ContainerAzureNodePool#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

