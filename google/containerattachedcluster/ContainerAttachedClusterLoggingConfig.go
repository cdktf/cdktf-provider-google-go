package containerattachedcluster


type ContainerAttachedClusterLoggingConfig struct {
	// component_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_attached_cluster#component_config ContainerAttachedCluster#component_config}
	ComponentConfig *ContainerAttachedClusterLoggingConfigComponentConfig `field:"optional" json:"componentConfig" yaml:"componentConfig"`
}

