package containerattachedcluster


type ContainerAttachedClusterLoggingConfig struct {
	// component_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_attached_cluster#component_config ContainerAttachedCluster#component_config}
	ComponentConfig *ContainerAttachedClusterLoggingConfigComponentConfig `field:"optional" json:"componentConfig" yaml:"componentConfig"`
}

