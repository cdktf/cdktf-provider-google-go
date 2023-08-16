package containerawsnodepool


type ContainerAwsNodePoolAutoscaling struct {
	// Maximum number of nodes in the NodePool. Must be >= min_node_count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_aws_node_pool#max_node_count ContainerAwsNodePool#max_node_count}
	MaxNodeCount *float64 `field:"required" json:"maxNodeCount" yaml:"maxNodeCount"`
	// Minimum number of nodes in the NodePool. Must be >= 1 and <= max_node_count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_aws_node_pool#min_node_count ContainerAwsNodePool#min_node_count}
	MinNodeCount *float64 `field:"required" json:"minNodeCount" yaml:"minNodeCount"`
}

