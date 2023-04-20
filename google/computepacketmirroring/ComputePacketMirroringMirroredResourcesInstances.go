package computepacketmirroring


type ComputePacketMirroringMirroredResourcesInstances struct {
	// The URL of the instances where this rule should be active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.62.1/docs/resources/compute_packet_mirroring#url ComputePacketMirroring#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

