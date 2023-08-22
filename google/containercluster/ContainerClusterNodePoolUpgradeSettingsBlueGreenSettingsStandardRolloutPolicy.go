package containercluster


type ContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy struct {
	// Number of blue nodes to drain in a batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_cluster#batch_node_count ContainerCluster#batch_node_count}
	BatchNodeCount *float64 `field:"optional" json:"batchNodeCount" yaml:"batchNodeCount"`
	// Percentage of the blue pool nodes to drain in a batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_cluster#batch_percentage ContainerCluster#batch_percentage}
	BatchPercentage *float64 `field:"optional" json:"batchPercentage" yaml:"batchPercentage"`
	// Soak time after each batch gets drained.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_cluster#batch_soak_duration ContainerCluster#batch_soak_duration}
	BatchSoakDuration *string `field:"optional" json:"batchSoakDuration" yaml:"batchSoakDuration"`
}

