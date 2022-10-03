package computepacketmirroring


type ComputePacketMirroringMirroredResourcesSubnetworks struct {
	// The URL of the subnetwork where this rule should be active.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_packet_mirroring#url ComputePacketMirroring#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

