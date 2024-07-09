// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computepacketmirroring


type ComputePacketMirroringCollectorIlb struct {
	// The URL of the forwarding rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/compute_packet_mirroring#url ComputePacketMirroring#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

