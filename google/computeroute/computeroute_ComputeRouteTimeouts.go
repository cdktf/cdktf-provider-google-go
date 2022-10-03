package computeroute


type ComputeRouteTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_route#create ComputeRoute#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_route#delete ComputeRoute#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

