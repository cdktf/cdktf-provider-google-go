package computeinstance


type ComputeInstanceBootDiskInitializeParams struct {
	// The image from which this disk was initialised.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance#image ComputeInstance#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
	// A set of key/value label pairs assigned to the disk.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance#labels ComputeInstance#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The size of the image in gigabytes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance#size ComputeInstance#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// The Google Compute Engine disk type. Such as pd-standard, pd-ssd or pd-balanced.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance#type ComputeInstance#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

