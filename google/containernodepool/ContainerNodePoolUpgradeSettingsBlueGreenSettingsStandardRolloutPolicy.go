package containernodepool


type ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy struct {
	// Number of blue nodes to drain in a batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_node_pool#batch_node_count ContainerNodePool#batch_node_count}
	BatchNodeCount *float64 `field:"optional" json:"batchNodeCount" yaml:"batchNodeCount"`
	// Percentage of the blue pool nodes to drain in a batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_node_pool#batch_percentage ContainerNodePool#batch_percentage}
	BatchPercentage *float64 `field:"optional" json:"batchPercentage" yaml:"batchPercentage"`
	// Soak time after each batch gets drained.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_node_pool#batch_soak_duration ContainerNodePool#batch_soak_duration}
	BatchSoakDuration *string `field:"optional" json:"batchSoakDuration" yaml:"batchSoakDuration"`
}

