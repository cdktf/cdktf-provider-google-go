package computepacketmirroring


type ComputePacketMirroringTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_packet_mirroring#create ComputePacketMirroring#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_packet_mirroring#delete ComputePacketMirroring#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_packet_mirroring#update ComputePacketMirroring#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

