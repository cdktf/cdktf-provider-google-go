// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ContainerNodePoolAutoscaling struct {
	// Maximum number of nodes in the NodePool. Must be >= min_node_count.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_node_pool#max_node_count ContainerNodePool#max_node_count}
	MaxNodeCount *float64 `field:"required" json:"maxNodeCount" yaml:"maxNodeCount"`
	// Minimum number of nodes in the NodePool. Must be >=0 and <= max_node_count.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_node_pool#min_node_count ContainerNodePool#min_node_count}
	MinNodeCount *float64 `field:"required" json:"minNodeCount" yaml:"minNodeCount"`
}

