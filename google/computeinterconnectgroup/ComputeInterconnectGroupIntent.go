// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinterconnectgroup


type ComputeInterconnectGroupIntent struct {
	// The reliability the user intends this group to be capable of, in terms of the Interconnect product SLAs.
	//
	// Possible values: ["PRODUCTION_NON_CRITICAL", "PRODUCTION_CRITICAL", "NO_SLA", "AVAILABILITY_SLA_UNSPECIFIED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/compute_interconnect_group#topology_capability ComputeInterconnectGroup#topology_capability}
	TopologyCapability *string `field:"optional" json:"topologyCapability" yaml:"topologyCapability"`
}

