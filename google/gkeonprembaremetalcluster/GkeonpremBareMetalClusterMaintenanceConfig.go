// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterMaintenanceConfig struct {
	// All IPv4 address from these ranges will be placed into maintenance mode.
	//
	// Nodes in maintenance mode will be cordoned and drained. When both of these
	// are true, the "baremetal.cluster.gke.io/maintenance" annotation will be set
	// on the node resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/gkeonprem_bare_metal_cluster#maintenance_address_cidr_blocks GkeonpremBareMetalCluster#maintenance_address_cidr_blocks}
	MaintenanceAddressCidrBlocks *[]*string `field:"required" json:"maintenanceAddressCidrBlocks" yaml:"maintenanceAddressCidrBlocks"`
}

