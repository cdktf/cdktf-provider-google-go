package containernodepool


type ContainerNodePoolNodeConfigGcfsConfig struct {
	// Whether or not GCFS is enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_node_pool#enabled ContainerNodePool#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

