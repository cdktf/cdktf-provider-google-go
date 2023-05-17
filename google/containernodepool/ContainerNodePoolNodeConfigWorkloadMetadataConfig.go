package containernodepool


type ContainerNodePoolNodeConfigWorkloadMetadataConfig struct {
	// Mode is the configuration for how to expose metadata to workloads running on the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_node_pool#mode ContainerNodePool#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

