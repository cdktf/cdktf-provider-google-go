// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computepacketmirroring


type ComputePacketMirroringMirroredResourcesSubnetworks struct {
	// The URL of the subnetwork where this rule should be active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/compute_packet_mirroring#url ComputePacketMirroring#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

