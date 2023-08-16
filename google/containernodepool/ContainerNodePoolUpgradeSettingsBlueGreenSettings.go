package containernodepool


type ContainerNodePoolUpgradeSettingsBlueGreenSettings struct {
	// standard_rollout_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_node_pool#standard_rollout_policy ContainerNodePool#standard_rollout_policy}
	StandardRolloutPolicy *ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy `field:"required" json:"standardRolloutPolicy" yaml:"standardRolloutPolicy"`
	// Time needed after draining entire blue pool. After this period, blue pool will be cleaned up.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_node_pool#node_pool_soak_duration ContainerNodePool#node_pool_soak_duration}
	NodePoolSoakDuration *string `field:"optional" json:"nodePoolSoakDuration" yaml:"nodePoolSoakDuration"`
}

