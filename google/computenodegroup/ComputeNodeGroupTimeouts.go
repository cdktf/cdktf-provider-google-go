package computenodegroup


type ComputeNodeGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_node_group#create ComputeNodeGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_node_group#delete ComputeNodeGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_node_group#update ComputeNodeGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
