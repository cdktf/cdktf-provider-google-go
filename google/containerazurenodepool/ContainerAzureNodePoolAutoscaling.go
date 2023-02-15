package containerazurenodepool


type ContainerAzureNodePoolAutoscaling struct {
	// Maximum number of nodes in the node pool. Must be >= min_node_count.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_azure_node_pool#max_node_count ContainerAzureNodePool#max_node_count}
	MaxNodeCount *float64 `field:"required" json:"maxNodeCount" yaml:"maxNodeCount"`
	// Minimum number of nodes in the node pool. Must be >= 1 and <= max_node_count.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_azure_node_pool#min_node_count ContainerAzureNodePool#min_node_count}
	MinNodeCount *float64 `field:"required" json:"minNodeCount" yaml:"minNodeCount"`
}

