package computesubnetwork


type ComputeSubnetworkTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_subnetwork#create ComputeSubnetwork#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_subnetwork#delete ComputeSubnetwork#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_subnetwork#update ComputeSubnetwork#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

