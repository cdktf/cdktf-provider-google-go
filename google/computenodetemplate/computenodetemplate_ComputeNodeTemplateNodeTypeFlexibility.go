package computenodetemplate


type ComputeNodeTemplateNodeTypeFlexibility struct {
	// Number of virtual CPUs to use.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_node_template#cpus ComputeNodeTemplate#cpus}
	Cpus *string `field:"optional" json:"cpus" yaml:"cpus"`
	// Physical memory available to the node, defined in MB.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_node_template#memory ComputeNodeTemplate#memory}
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}

