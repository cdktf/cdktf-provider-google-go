package computepacketmirroring


type ComputePacketMirroringMirroredResources struct {
	// instances block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_packet_mirroring#instances ComputePacketMirroring#instances}
	Instances interface{} `field:"optional" json:"instances" yaml:"instances"`
	// subnetworks block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_packet_mirroring#subnetworks ComputePacketMirroring#subnetworks}
	Subnetworks interface{} `field:"optional" json:"subnetworks" yaml:"subnetworks"`
	// All instances with these tags will be mirrored.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_packet_mirroring#tags ComputePacketMirroring#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

