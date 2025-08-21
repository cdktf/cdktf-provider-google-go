// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionbackendservice


type ComputeRegionBackendServiceHaPolicyLeaderNetworkEndpoint struct {
	// The name of the VM instance of the leader network endpoint.
	//
	// The instance must
	// already be attached to the NEG specified in the haPolicy.leader.backendGroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/compute_region_backend_service#instance ComputeRegionBackendService#instance}
	Instance *string `field:"optional" json:"instance" yaml:"instance"`
}

