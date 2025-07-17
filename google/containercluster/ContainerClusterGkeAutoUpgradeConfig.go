// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterGkeAutoUpgradeConfig struct {
	// The selected auto-upgrade patch type.
	//
	// Accepted values are:
	// * ACCELERATED: Upgrades to the latest available patch version in a given minor and release channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/container_cluster#patch_mode ContainerCluster#patch_mode}
	PatchMode *string `field:"required" json:"patchMode" yaml:"patchMode"`
}

