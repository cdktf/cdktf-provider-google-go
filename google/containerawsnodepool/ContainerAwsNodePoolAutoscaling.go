package containerawsnodepool


type ContainerAwsNodePoolAutoscaling struct {
	// Maximum number of nodes in the NodePool. Must be >= min_node_count.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_aws_node_pool#max_node_count ContainerAwsNodePool#max_node_count}
	MaxNodeCount *float64 `field:"required" json:"maxNodeCount" yaml:"maxNodeCount"`
	// Minimum number of nodes in the NodePool. Must be >= 1 and <= max_node_count.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_aws_node_pool#min_node_count ContainerAwsNodePool#min_node_count}
	MinNodeCount *float64 `field:"required" json:"minNodeCount" yaml:"minNodeCount"`
}

