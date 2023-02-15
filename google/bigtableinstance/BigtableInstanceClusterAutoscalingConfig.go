package bigtableinstance


type BigtableInstanceClusterAutoscalingConfig struct {
	// The target CPU utilization for autoscaling. Value must be between 10 and 80.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_instance#cpu_target BigtableInstance#cpu_target}
	CpuTarget *float64 `field:"required" json:"cpuTarget" yaml:"cpuTarget"`
	// The maximum number of nodes for autoscaling.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_instance#max_nodes BigtableInstance#max_nodes}
	MaxNodes *float64 `field:"required" json:"maxNodes" yaml:"maxNodes"`
	// The minimum number of nodes for autoscaling.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_instance#min_nodes BigtableInstance#min_nodes}
	MinNodes *float64 `field:"required" json:"minNodes" yaml:"minNodes"`
	// The target storage utilization for autoscaling, in GB, for each node in a cluster.
	//
	// This number is limited between 2560 (2.5TiB) and 5120 (5TiB) for a SSD cluster and between 8192 (8TiB) and 16384 (16 TiB) for an HDD cluster. If not set, whatever is already set for the cluster will not change, or if the cluster is just being created, it will use the default value of 2560 for SSD clusters and 8192 for HDD clusters.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_instance#storage_target BigtableInstance#storage_target}
	StorageTarget *float64 `field:"optional" json:"storageTarget" yaml:"storageTarget"`
}

