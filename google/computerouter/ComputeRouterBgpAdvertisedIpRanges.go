package computerouter


type ComputeRouterBgpAdvertisedIpRanges struct {
	// The IP range to advertise. The value must be a CIDR-formatted string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.0/docs/resources/compute_router#range ComputeRouter#range}
	Range *string `field:"required" json:"range" yaml:"range"`
	// User-specified description for the IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.0/docs/resources/compute_router#description ComputeRouter#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

