package computerouterpeer


type ComputeRouterPeerAdvertisedIpRanges struct {
	// The IP range to advertise. The value must be a CIDR-formatted string.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_router_peer#range ComputeRouterPeer#range}
	Range *string `field:"required" json:"range" yaml:"range"`
	// User-specified description for the IP range.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_router_peer#description ComputeRouterPeer#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

