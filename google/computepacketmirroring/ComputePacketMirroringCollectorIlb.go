package computepacketmirroring


type ComputePacketMirroringCollectorIlb struct {
	// The URL of the forwarding rule.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_packet_mirroring#url ComputePacketMirroring#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

