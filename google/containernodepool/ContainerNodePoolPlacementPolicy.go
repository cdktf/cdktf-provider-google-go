package containernodepool


type ContainerNodePoolPlacementPolicy struct {
	// Type defines the type of placement policy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_node_pool#type ContainerNodePool#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

