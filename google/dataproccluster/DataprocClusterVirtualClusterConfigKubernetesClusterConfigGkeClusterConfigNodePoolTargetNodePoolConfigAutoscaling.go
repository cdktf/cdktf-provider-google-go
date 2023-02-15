package dataproccluster


type DataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling struct {
	// The maximum number of nodes in the node pool. Must be >= minNodeCount, and must be > 0.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_cluster#max_node_count DataprocCluster#max_node_count}
	MaxNodeCount *float64 `field:"optional" json:"maxNodeCount" yaml:"maxNodeCount"`
	// The minimum number of nodes in the node pool. Must be >= 0 and <= maxNodeCount.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_cluster#min_node_count DataprocCluster#min_node_count}
	MinNodeCount *float64 `field:"optional" json:"minNodeCount" yaml:"minNodeCount"`
}

