package computedisk


type ComputeDiskTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_disk#create ComputeDisk#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_disk#delete ComputeDisk#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_disk#update ComputeDisk#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

