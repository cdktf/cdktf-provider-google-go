package containernodepool


type ContainerNodePoolTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_node_pool#create ContainerNodePool#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_node_pool#delete ContainerNodePool#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_node_pool#update ContainerNodePool#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

