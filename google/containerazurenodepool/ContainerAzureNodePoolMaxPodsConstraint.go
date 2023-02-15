package containerazurenodepool


type ContainerAzureNodePoolMaxPodsConstraint struct {
	// The maximum number of pods to schedule on a single node.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_azure_node_pool#max_pods_per_node ContainerAzureNodePool#max_pods_per_node}
	MaxPodsPerNode *float64 `field:"required" json:"maxPodsPerNode" yaml:"maxPodsPerNode"`
}

