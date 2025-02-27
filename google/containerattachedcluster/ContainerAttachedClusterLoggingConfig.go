// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerattachedcluster


type ContainerAttachedClusterLoggingConfig struct {
	// component_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/container_attached_cluster#component_config ContainerAttachedCluster#component_config}
	ComponentConfig *ContainerAttachedClusterLoggingConfigComponentConfig `field:"optional" json:"componentConfig" yaml:"componentConfig"`
}

