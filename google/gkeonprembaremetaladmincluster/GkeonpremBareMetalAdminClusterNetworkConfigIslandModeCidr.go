// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetaladmincluster


type GkeonpremBareMetalAdminClusterNetworkConfigIslandModeCidr struct {
	// All pods in the cluster are assigned an RFC1918 IPv4 address from these ranges.
	//
	// This field cannot be changed after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.17.0/docs/resources/gkeonprem_bare_metal_admin_cluster#pod_address_cidr_blocks GkeonpremBareMetalAdminCluster#pod_address_cidr_blocks}
	PodAddressCidrBlocks *[]*string `field:"required" json:"podAddressCidrBlocks" yaml:"podAddressCidrBlocks"`
	// All services in the cluster are assigned an RFC1918 IPv4 address from these ranges.
	//
	// This field cannot be changed after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.17.0/docs/resources/gkeonprem_bare_metal_admin_cluster#service_address_cidr_blocks GkeonpremBareMetalAdminCluster#service_address_cidr_blocks}
	ServiceAddressCidrBlocks *[]*string `field:"required" json:"serviceAddressCidrBlocks" yaml:"serviceAddressCidrBlocks"`
}

