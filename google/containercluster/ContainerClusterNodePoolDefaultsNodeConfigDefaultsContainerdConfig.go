// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfig struct {
	// private_registry_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/container_cluster#private_registry_access_config ContainerCluster#private_registry_access_config}
	PrivateRegistryAccessConfig *ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigPrivateRegistryAccessConfig `field:"optional" json:"privateRegistryAccessConfig" yaml:"privateRegistryAccessConfig"`
}

