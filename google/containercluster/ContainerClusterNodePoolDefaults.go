package containercluster


type ContainerClusterNodePoolDefaults struct {
	// node_config_defaults block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#node_config_defaults ContainerCluster#node_config_defaults}
	NodeConfigDefaults *ContainerClusterNodePoolDefaultsNodeConfigDefaults `field:"optional" json:"nodeConfigDefaults" yaml:"nodeConfigDefaults"`
}

