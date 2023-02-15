package computeperinstanceconfig


type ComputePerInstanceConfigPreservedState struct {
	// disk block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_per_instance_config#disk ComputePerInstanceConfig#disk}
	Disk interface{} `field:"optional" json:"disk" yaml:"disk"`
	// Preserved metadata defined for this instance. This is a list of key->value pairs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_per_instance_config#metadata ComputePerInstanceConfig#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
}

