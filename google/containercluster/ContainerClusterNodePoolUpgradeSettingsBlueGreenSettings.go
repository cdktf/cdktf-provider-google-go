package containercluster


type ContainerClusterNodePoolUpgradeSettingsBlueGreenSettings struct {
	// standard_rollout_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_cluster#standard_rollout_policy ContainerCluster#standard_rollout_policy}
	StandardRolloutPolicy *ContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy `field:"required" json:"standardRolloutPolicy" yaml:"standardRolloutPolicy"`
	// Time needed after draining entire blue pool. After this period, blue pool will be cleaned up.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_cluster#node_pool_soak_duration ContainerCluster#node_pool_soak_duration}
	NodePoolSoakDuration *string `field:"optional" json:"nodePoolSoakDuration" yaml:"nodePoolSoakDuration"`
}

