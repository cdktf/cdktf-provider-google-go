package containernodepool


type ContainerNodePoolNodeConfigTaint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.64.0/docs/resources/container_node_pool#effect ContainerNodePool#effect}.
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.64.0/docs/resources/container_node_pool#key ContainerNodePool#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.64.0/docs/resources/container_node_pool#value ContainerNodePool#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

