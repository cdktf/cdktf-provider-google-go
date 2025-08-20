// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionbackendservice


type ComputeRegionBackendServiceHaPolicyLeader struct {
	// A fully-qualified URL of the zonal Network Endpoint Group (NEG) that the leader is attached to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/compute_region_backend_service#backend_group ComputeRegionBackendService#backend_group}
	BackendGroup *string `field:"optional" json:"backendGroup" yaml:"backendGroup"`
	// network_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/compute_region_backend_service#network_endpoint ComputeRegionBackendService#network_endpoint}
	NetworkEndpoint *ComputeRegionBackendServiceHaPolicyLeaderNetworkEndpoint `field:"optional" json:"networkEndpoint" yaml:"networkEndpoint"`
}

