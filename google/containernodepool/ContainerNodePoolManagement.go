// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containernodepool


type ContainerNodePoolManagement struct {
	// Whether the nodes will be automatically repaired. Enabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/container_node_pool#auto_repair ContainerNodePool#auto_repair}
	AutoRepair interface{} `field:"optional" json:"autoRepair" yaml:"autoRepair"`
	// Whether the nodes will be automatically upgraded. Enabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/container_node_pool#auto_upgrade ContainerNodePool#auto_upgrade}
	AutoUpgrade interface{} `field:"optional" json:"autoUpgrade" yaml:"autoUpgrade"`
}

