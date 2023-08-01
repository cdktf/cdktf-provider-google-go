package containernodepool


type ContainerNodePoolNodeConfigSoleTenantConfigNodeAffinity struct {
	// .
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_node_pool#key ContainerNodePool#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// .
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_node_pool#operator ContainerNodePool#operator}
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// .
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_node_pool#values ContainerNodePool#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

