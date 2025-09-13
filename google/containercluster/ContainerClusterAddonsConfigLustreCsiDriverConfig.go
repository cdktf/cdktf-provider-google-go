// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterAddonsConfigLustreCsiDriverConfig struct {
	// Whether the Lustre CSI driver is enabled for this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// If set to true, the Lustre CSI driver will initialize LNet (the virtual network layer for Lustre kernel module) using port 6988.
	//
	// This flag is required to workaround a port conflict with the gke-metadata-server on GKE nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_cluster#enable_legacy_lustre_port ContainerCluster#enable_legacy_lustre_port}
	EnableLegacyLustrePort interface{} `field:"optional" json:"enableLegacyLustrePort" yaml:"enableLegacyLustrePort"`
}

