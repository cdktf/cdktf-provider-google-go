package computepacketmirroring


type ComputePacketMirroringFilter struct {
	// IP CIDR ranges that apply as a filter on the source (ingress) or destination (egress) IP in the IP header.
	//
	// Only IPv4 is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_packet_mirroring#cidr_ranges ComputePacketMirroring#cidr_ranges}
	CidrRanges *[]*string `field:"optional" json:"cidrRanges" yaml:"cidrRanges"`
	// Direction of traffic to mirror. Default value: "BOTH" Possible values: ["INGRESS", "EGRESS", "BOTH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_packet_mirroring#direction ComputePacketMirroring#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Possible IP protocols including tcp, udp, icmp and esp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_packet_mirroring#ip_protocols ComputePacketMirroring#ip_protocols}
	IpProtocols *[]*string `field:"optional" json:"ipProtocols" yaml:"ipProtocols"`
}

